fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test ./...
build:
	go build ./cmd/gocore
run:
	go run ./cmd/gocore
up:
	docker compose up -d
down:
	docker compose down