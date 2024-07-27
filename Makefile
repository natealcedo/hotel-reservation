.PHONY: build clean test run install kill

APP_NAME = app
BIN = bin/$(APP_NAME)

dev:
	@pnpm pm2 start ecosystem.config.js --no-daemon

kill:
	@pnpm pm2 kill

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