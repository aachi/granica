#!/bin/sh

# Clear
clear

# Free ports
killall -9 main
killall -9 runner-build

# Build
./scripts/build.sh

# Set environment variables

# Repo
export REPO_TYPE="mongodb"
export REPO_HOST="localhost"
export REPO_PORT=27017
export REPO_DB="granica"
export REPO_USER="granica"
export REPO_PASSWORD="granica"

# Broker
export BROKER_TYPE="rabbitmq"
export RABBITMQ_HOST="localhost"
export RABBITMQ_PORT="5672"
export RABBITMQ_USER="granica"
export RABBITMQ_PASSWORD="granica"
export RABBITMQ_EXCHANGE="default"
export RABBITMQ_QUEUE="main"

# Cache
export CACHE_TYPE="redis"
export REDIS_HOST="localhost"
export REDIS_PORT=6379
export REDIS_DB=0
export REDIS_USER="granica"
export REDIS_PASSWORD="granica"

# Start
go run main.go

# Ref.: Fresh - https://github.com/gravityblast/fresh
# go get github.com/pilu/fresh
# Start with auto-reload
# fresh

