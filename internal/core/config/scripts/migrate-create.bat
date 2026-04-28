@echo off
if "%seq%"=="" (
    echo ERROR: seq is required. Usage: make migrate-create seq=001
    exit /b 1
)
docker compose run --rm polls-app-postgres-migrate create -ext sql -dir /migrations -seq "%seq%"
