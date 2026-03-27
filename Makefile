run:
	go run ./cmd/api

tidy:
	go mod tidy

fmt:
	go fmt ./...

test:
	go test ./...

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build

dev-down:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml down