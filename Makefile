# Define the environment (default is 'dev')
ENV ?= dev

# these are the default values
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:root1234@tcp(127.0.0.1:33336)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

gen-proto:
	cd proto && sh gen.sh
	
gen-swagger:
	swag init -g ./cmd/ecommerce/main.go -o ./cmd/swag/docs

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

# add one table to database
up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
	
# create a new migration
create-migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

up-gen:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

down-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

