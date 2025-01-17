#!/bin/bash

DB_USER=$1
DB_PASSWORD=$2
DB_HOST=$3
DB_PORT=$4
DB_NAME=$5
DB_SSLMODE=$6

SEED_FILE="$(dirname "$0")/seeds.sql"

if [ ! -f  "$SEED_FILE" ]; then
  echo "File seeds.sql not found!"
  exit 1
fi

echo "Start seeding..."

if [ "$DB_SSLMODE" == "true" ]; then
  MYSQL_CMD="mysql -u $DB_USER -p$DB_PASSWORD -h $DB_HOST -P $DB_PORT --ssl-mode=REQUIRED $DB_NAME"
else
  MYSQL_CMD="mysql -u $DB_USER -p$DB_PASSWORD -h $DB_HOST -P $DB_PORT $DB_NAME"
fi

$MYSQL_CMD <  "$SEED_FILE"

if [ $? -eq 0 ]; then
  echo "Seed data success!"
else
  echo "Something wrong when seeding data..."
  exit 1
fi
