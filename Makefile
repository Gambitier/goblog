DB_HOST=$(shell jq -r '.database.host' config.json)
DB_PORT=$(shell jq -r '.database.port' config.json)
DB_USER=$(shell jq -r '.database.user' config.json)
DB_PASS=$(shell jq -r '.database.password' config.json)
DB_NAME=$(shell jq -r '.database.dbname' config.json)
DB_SSL=$(shell jq -r '.database.sslmode' config.json)

DB_URL=postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)

.PHONY: server migrate-up migrate-down sqlc

server:
	go run main.go

sqlc:
	cd database && sqlc generate

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down 