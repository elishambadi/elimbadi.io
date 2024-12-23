# Variables
APP_NAME := elimbadi.io
CMD_DIR := cmd/server
PKG := github.com/elishambadi/elimbadi.io

.PHONY: all build run test clearn tidy

all: build

## Builds the application
build:
    @echo "==> Building the application..."
    go build -o  bin/$(APP_NAME) $(CMD_DIR)/main.go
