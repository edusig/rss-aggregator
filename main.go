package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// type apiConfig struct {
// 	jwtSecret string
// }

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	// jwtSecret := os.Getenv("JST_SECRET")

	// apiCfg := apiConfig{
	// 	jwtSecret: jwtSecret,
	// }

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
	r.Mount("/v1", apiV1Router)

	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
