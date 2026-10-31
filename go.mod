module github.com/edusig/rss-aggregator

go 1.24

require (
	github.com/go-chi/chi/v5 v5.4.0
	github.com/go-chi/cors v1.2.1
	github.com/google/uuid v1.3.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	internal/auth v1.0.0
	internal/database v1.0.0
)

replace internal/database => ./internal/database

replace internal/auth => ./internal/auth
