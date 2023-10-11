run-build:
	make run-docker-db &
	go build -o out && ./out

run-docker-db:
	docker start postgres

run-sql-generate:
	sqlc generate

create-docker-db:
	docker run -v ${DOCKER_POSTGRES_DATA}:/var/lib/postgresql/data -d -e POSTGRES_PASSWORD=foobarbin -e POSTGRES_USER=postgres -e PGDATA=/var/lib/postgresql/data -p 9999:5432 --name postgres postgres:15

db-migrate-up:
	cd sql/schema && goose postgres postgres://postgres:foobarbin@localhost:9999/rss-aggregator up

db-migrate-down:
	cd sql/schema && goose postgres postgres://postgres:foobarbin@localhost:9999/rss-aggregator down

db-reset:
	cd sql/schema && goose postgres postgres://postgres:foobarbin@localhost:9999/rss-aggregator reset

db-status:
	cd sql/schema && goose postgres postgres://postgres:foobarbin@localhost:9999/rss-aggregator status

build-docker-image:
	GOOS=linux GOARCH=amd64 go build
	docker build . -t edusig/rss-aggregator:latest -t edusig/rss-aggregator:${VERSION}