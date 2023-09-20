#!/bin/sh

set -e

echo "Running database migrations..."
migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting server..."
exec "$@"
