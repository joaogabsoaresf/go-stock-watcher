.PHONY: default run build test docs clean
# Variables
APP_NAME=mygocrm

# Tasks
default: run-with-docs

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
run-with-docs:
	@swag init
	@go run main.go
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs