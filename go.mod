module github.com/edusig/rss-aggregator

go 1.21.1

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	internal/database v1.0.0
	internal/auth v1.0.0
	github.com/google/uuid v1.3.1
)

replace internal/database => ./internal/database

replace internal/auth => ./internal/auth
