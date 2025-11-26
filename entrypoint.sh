#!/bin/sh
set -e

echo "Waiting for database to be ready..."
# A simple wait loop, you might want a more robust solution like wait-for-it.sh
# This requires netcat (nc) to be installed in the final image.
# The host and port should be extracted from DATABASE_URL or passed as separate env vars.
# For now, this is a placeholder. A proper implementation would parse the DB_URL.
# Let's assume a simple sleep for now to avoid adding dependencies.
sleep 5

echo "Running database migrations..."
/app/migrate -path /app/database/migrations -database "$URL_POSTGRES_SENSORS" up

echo "Starting application..."
exec /app/main