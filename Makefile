include .env
export

env-up:
	@docker compose up -d polls-app-postgres

env-down:
	@docker compose down

env-cleanup:
	@cmd\scripts\env-cleanup.bat

migrate-create:
	@cmd\scripts\migrate-create.bat

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@cmd\scripts\migrate-action.bat