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