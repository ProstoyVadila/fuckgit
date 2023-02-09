
run_front:
	@echo "Starting frontend..."
	@cd frontend && npm start

run_back:
	@echo "Starting backend..."
	@cd backend && go run cmd/main.go

run: run_front run_back

