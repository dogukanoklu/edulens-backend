#!/bin/bash

# Load environment variables
set -a
source ../../cmd/.env
set +a

# MySQL command to set up the database
echo "Setting up the database..."
mysql -u "$DB_USER" -p"$DB_PASSWORD" -h "$DB_HOST" -P "$DB_PORT" "$DB_NAME" < schema.sql
echo "Database setup completed."
