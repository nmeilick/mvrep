# Makefile for building mvrep for multiple platforms

BINARY_NAME=mvrep
BUILD_DIR=bin

PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 windows/amd64

# Get the current git tag or use "dev" as fallback, remove "v" prefix if present
VERSION := $(shell git describe --tags --always --dirty="-dev" | sed 's/^v//')

all: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-$(word 1, $(subst /, ,$@))-$(word 2, $(subst /, ,$@)) main.go

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all clean $(PLATFORMS)
