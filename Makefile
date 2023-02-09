 BACK_BINARY=webApp

up:
	@echo "Starting docker containers..."
	docker compose up -d
	@echo "Done!"

down:
	@echo "Stopping docker containers..."
	docker compose down
	@echo "Done!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Starting docker containers with build..."
	docker compose up -d --build
	@echo "Done!"

run_front:
	@echo "Starting frontend locally..."
	@cd frontend && npm start

run_back:
	@echo "Starting backend locally..."
	@cd backend && ADDRESS=localhost:8080 go run ./cmd/api

run: run_front run_back

build_front:
	@echo "Building frontend..."
	@cd frontend && npm run build

build_back: fieldalignment
	@echo "Building backend..."
	@cd backend && env GOOS=linux CGO_ENABLED=0 go build -o ${BACK_BINARY} ./cmd/api

build: build_front build_back

fieldalignment:
	@echo "Checking fieldalignment..."
	@cd backend && fieldalignment -fix ./...

.PHONY: run_front run_back run build_front build_back build
