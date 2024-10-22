docker-build:
	docker-compose up -d --build
	docker-compose ps

docker-stop:
	docker-compose down

docker-up:
	docker compose up -d

local:
	go run ./cmd/ecommerce

