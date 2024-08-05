# Makefile for building mvrep for multiple platforms

BINARY_NAME=mvrep
BUILD_DIR=bin

PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 windows/amd64

all: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(word 1, $(subst /, ,$@)) GOARCH=$(word 2, $(subst /, ,$@)) go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(word 1, $(subst /, ,$@))-$(word 2, $(subst /, ,$@)) main.go

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all clean $(PLATFORMS)
