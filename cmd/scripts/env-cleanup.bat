@echo off
echo WARNING: This will delete all database data and the 'out' directory.
set /p "confirm=Are you sure? [y/N]: "
if /I "%confirm%"=="y" (
    docker compose down -v --remove-orphans
    if exist "out" rmdir /s /q out
    echo Cleanup complete.
) else (
    echo Cleanup cancelled.
)
