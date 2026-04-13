DB_PATH=./db

include .env
export

DB_URL=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

migrate:
	migrate -path ./db -database "$(DB_URL)" up

# migrate:
# 	migrate -path /Users/mozhno/Documents/golang/finance/db -database 'postgres://postgres:qwerty@0.0.0.0:5442/postgres?sslmode=disable' up
	
tidy:
	go mod tidy && go fmt ./...