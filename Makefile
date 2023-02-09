up:
	@echo "Starting docker containers..."
	docker compose up -d

down:
	@echo "Stopping docker containers..."
	docker compose down

run_front:
	@echo "Starting frontend..."
	@cd frontend && npm start

run_back:
	@echo "Starting backend..."
	@cd backend && go run cmd/api/main.go

run: run_front run_back

build_front:
	@echo "Building frontend..."
	@cd frontend && npm run build

build_back:
	@echo "Building backend..."
	@cd backend && go build -o bin/api cmd/api/main.go

build: build_front build_back

.PHONY: run_front run_back run build_front build_back build
