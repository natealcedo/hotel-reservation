.PHONY: build clean test run install lint

APP_NAME = app
BIN = bin/$(APP_NAME)

run: build
	@./$(BIN)

install:
	@go mod tidy

build:
	@go build -o $(BIN)

clean:
	@rm -rf bin/*

test:
	@go test -v ./... -count=1 # -count=1 to avoid caching

lint:
	@if gofmt -l . | read; then \
		gofmt -d .; \
		exit 1; \
	fi
