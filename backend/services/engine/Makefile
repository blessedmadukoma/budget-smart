include .env

build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -cover -v ./...

run: build
	@./bin/ecom

air:
	@air

dc_up:
	@docker compose up -d

dc_down:
	@docker compose down -v

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@echo "Running migrations up..."
	@go run cmd/migrate/main.go up
	@echo "Migrate Up Complete..."

migrate-down:
	@echo "Running migrations down..."
	@go run cmd/migrate/main.go down
	@echo "Migrate Down Complete..."
