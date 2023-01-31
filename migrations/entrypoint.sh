#!/bin/bash
export GOOSE_DRIVER="postgres"

export GOOSE_DBSTRING="host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=$DB_SSL"

goose -dir ./db up
