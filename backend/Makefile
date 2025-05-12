.PHONY: dev-up dev-down build-engine build-ml-analytics build-notification \
build-all test-engine test-ml-analytics test-notification test-all \
logs logs-engine logs-ml-analytics logs-notification db-shell

# Development environment
dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

# Building services
build-engine:
	docker-compose build engine

build-ml-analytics:
	docker-compose build ml-analytics

build-notification:
	docker-compose build notification

build-all: build-engine build-ml-analytics build-notification

# Testing
test-engine:
	cd services/engine && go test ./...

test-ml-analytics:
	cd services/ml-analytics && python -m pytest

test-notification:
	cd services/notification && python -m pytest

test-all: test-engine test-ml-analytics test-notification

# Logs
logs:
	docker-compose logs -f

logs-engine:
	docker-compose logs -f engine

logs-ml-analytics:
	docker-compose logs -f ml-analytics

logs-notification:
	docker-compose logs -f notification

# Database
db-shell:
	docker-compose exec postgres psql -U budgetsmart -d budgetsmart

# RabbitMQ
rabbitmq-shell:
	docker-compose exec rabbitmq rabbitmqctl

# Migration
migrate:
	docker-compose exec engine go run ./cmd/migrations

seed-data:
	docker-compose exec engine go run ./cmd/seed

# Local development setup
setup-local:
	cp .env.example .env
	make dev-up
	make migrate
	make seed-data
	@echo "Local development environment is ready!"