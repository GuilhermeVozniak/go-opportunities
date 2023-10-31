.PHONEY: default run build test docs clean
# Variables
APP_NAME=goopportunities


#Tasks
default: run-with-docs

run:
	@go run -tags=sonic main.go
run-with-docs:
	@swag init
	@go run -tags=sonic main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs