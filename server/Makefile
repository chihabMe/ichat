.PHONY: run migrate build

run:
	nodemon --exec go run ./cmd/main.go --signal SIGTERM
migrate:
	go run ./cmd/migrate/migrate.go
build:
	go build -o ./build/bin ./cmd/main.go