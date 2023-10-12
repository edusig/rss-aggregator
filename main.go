package main

import (
	"database/sql"
	"internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbConnURL := os.Getenv("DATABASE_CONNECTION_URL")

	db, err := sql.Open("postgres", dbConnURL)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	apiV1Router := chi.NewRouter()
	apiV1Router.Get("/readiness", readinessHandler)
	apiV1Router.Get("/err", errorHandler)

	apiV1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.userGetHandler))
	apiV1Router.Post("/users", apiCfg.userCreateHandler)

	apiV1Router.Get("/feeds", apiCfg.getAllFeedsHandler)
	apiV1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.createFeedHandler))

	apiV1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.getFeedFollowsHandler))
	apiV1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.followFeedHandler))
	apiV1Router.Delete("/feed_follows", apiCfg.middlewareAuth(apiCfg.unfollowFeedHandler))

	apiV1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.getPostsByUser))

	r.Mount("/v1", apiV1Router)

	go FeedWorker(dbQueries)

	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
