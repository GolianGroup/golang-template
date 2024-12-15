# Variables
DEV_CONFIG_FILE := config/config.dev.yml
DEV_COMPOSE_FILE := build/docker-compose/docker-compose-dev.yml

# Extract values from config.dev.yml using yq (you need to install yq first)
DEV_DB_PASSWORD := $(shell yq e '.db.password' $(DEV_CONFIG_FILE))
DEV_DB_NAME := $(shell yq e '.db.dbname' $(DEV_CONFIG_FILE))
DEV_DB_USER := $(shell yq e '.db.user' $(DEV_CONFIG_FILE))
DEV_ARANGO_PASSWORD := $(shell yq e '.arango.password' $(DEV_CONFIG_FILE))
DEV_REDIS_PASSWORD := $(shell yq e '.redis.password' $(DEV_CONFIG_FILE))

# Development commands
.PHONY: dev
dev:
	@echo "Starting development environment..."
	DB_PASSWORD=$(DEV_DB_PASSWORD) \
	DB_NAME=$(DEV_DB_NAME) \
	DB_USER=$(DEV_DB_USER) \
	ARANGO_PASSWORD=$(DEV_ARANGO_PASSWORD) \
	REDIS_PASSWORD=$(DEV_REDIS_PASSWORD) \
	docker-compose -f $(DEV_COMPOSE_FILE) up -d

.PHONY: dev-down
dev-down:
	docker-compose -f $(DEV_COMPOSE_FILE) down

.PHONY: dev-logs
dev-logs:
	docker-compose -f $(DEV_COMPOSE_FILE) logs -f

.PHONY: dev-ps
dev-ps:
	docker-compose -f $(DEV_COMPOSE_FILE) ps 

.PHONY: gen-proto
gen-proto:
	protoc \
	--proto_path=./proto \
	--go_out=./proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_opt=require_unimplemented_servers=false \
	--experimental_allow_proto3_optional \
	--go-grpc_out=./proto \
	./proto/*.proto
