.PHONY: build run test lint up down db-up db-in mig-cr clean postgres createdb dropdb migrateup migratedown sqlc

up: db-up sl migrate
	DOCKER_BUILDKIT=1 docker-compose up --build server

down:
	docker-compose down

sl:
	sleep 3

run: up

# TODO dockerイメージをビルドするように書き換える
build:
	go build

testcov:
	go test -v -cover ./...

test:
	./go.test.sh

lint:
	golangci-lint run

clean:
	docker system prune --volumes -f

db-up:
	DOCKER_BUILDKIT=1 docker-compose up --build -d db

db-in:
	docker-compose exec db bash

mig-cr:
	migrate create -ext sql -dir db/migrations -seq ${name}

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrate:
	migrate -path db/migrations -database 'postgres://test:test@localhost:5432/template?sslmode=disable' -verbose up

migrateup:
	migrate -path db/migrations -database ${POSTGRESQL_URL} -verbose up

migratedown:
	migrate -path db/migrations -database ${POSTGRESQL_URL} -verbose down

sqlcgen:
	sqlc generate
