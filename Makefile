docker-build:
	docker-compose up --build

docker-stop:
	docker-compose down

docker-up:
	docker-compose up

local:
	go run ./cmd/ecommerce

