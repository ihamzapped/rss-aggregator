# Load environment variables from .env file
include .env

build: 
	@go build -o ./bin/rssagg ./cmd

run: build
	@./bin/rssagg

gendb:
	@sqlc generate

migratedb:
	@cd ./sql/schema; goose postgres $(DB_URL) up

dropdb:
	@cd ./sql/schema; goose postgres $(DB_URL) reset

refreshdb: dropdb migratedb gendb
