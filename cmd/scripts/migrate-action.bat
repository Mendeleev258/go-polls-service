@echo off
if "%action%"=="" (
    echo ERROR: action is required. Usage: make migrate-action action=up [or down]
    exit /b 1
)
docker compose run --rm polls-app-postgres-migrate ^
    -path /migrations ^
    -database "postgres://%POSTGRES_USER%:%POSTGRES_PASSWORD%@polls-app-postgres:5432/%POSTGRES_DB%?sslmode=disable" ^
    "%action%"
