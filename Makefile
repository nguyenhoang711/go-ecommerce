# Define the environment (default is 'dev')
ENV ?= dev

# these are the default values
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:123456@tcp(127.0.0.1:3306)/ecom_shop"
GOOSE_MIGRATION_DIR ?= sql/schema

docker-build:
	docker-compose up --build

docker-stop:
	docker-compose down

docker-up:
	docker-compose up

local:
	go run ./cmd/ecommerce

dev:
	@echo "Running with environment: $(ENV)"
	ENV=$(ENV) go run ./cmd/ecommerce

up-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

down-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate

