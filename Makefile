include internal/config/.env
export

env-up:
	@docker compose up -d polls-app-postgres

env-down:
	@docker compose down

env-cleanup:
	@internal\config\scripts\env-cleanup.bat

migrate-create:
	@internal\config\scripts\migrate-create.bat

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@internal\config\scripts\migrate-action.bat

pollsapp-run:
	@go run cmd/server/main.go