.PHONY: env-up env-down env-cleanup migrate-create migrate-up migrate-down migrate-action swagger-docs pollsapp-run

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

swagger-docs:
	C:\Users\jrosl\go\bin\swag.exe init -g cmd/server/main.go -o docs

pollsapp-run:
	@go run cmd/server/main.go