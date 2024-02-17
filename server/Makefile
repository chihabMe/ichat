.PHONY: run migrate

run:
	nodemon --exec go run ./cmd/main.go --signal SIGTERM
migrate:
	go run ./cmd/migrate/migrate.go