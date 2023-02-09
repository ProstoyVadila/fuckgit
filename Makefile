 BACK_BINARY=webApp

up:
	@echo "Starting docker containers..."
	docker compose up -d

down:
	@echo "Stopping docker containers..."
	docker compose down

up_build: build_back
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Starting docker containers with build..."
	docker compose up -d --build

run_front:
	@echo "Starting frontend locally..."
	@cd frontend && npm start

run_back:
	@echo "Starting backend locally..."
	@cd backend && go run ./cmd/api

run: run_front run_back

build_front:
	@echo "Building frontend..."
	@cd frontend && npm run build

build_back:
	@echo "Building backend..."
	@cd backend && env GOOS=linux CGO_ENABLED=0 go build -o ${BACK_BINARY} ./cmd/api

build: build_front build_back

.PHONY: run_front run_back run build_front build_back build
