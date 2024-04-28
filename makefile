.PHONY: default run build test docs clean
# Variables
APP_NAME=mygocrm

# Tasks
default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go