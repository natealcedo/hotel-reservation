.PHONY: build clean test run install lint fix

APP_NAME = app
BIN = bin/$(APP_NAME)

run:
	@air

install:
	@go mod tidy

build:
	@go build -o $(BIN)

clean:
	@rm -rf bin/*

test:
	@go test -v ./... -count=1 # -count=1 to avoid caching

lint:
	@gofmt -l .

fix:
	@gofmt -w .

seed:
	@go run scripts/main.go