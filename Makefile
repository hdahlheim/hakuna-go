BINARY_NAME=hakuna-go

all: build-macos-amd64 build-macos-arm64 build-linux-amd64 build-linux-arm64

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/$(BINARY_NAME)-macos-amd64
	md5sum ./dist/$(BINARY_NAME)-macos-amd64 | head -c 32 | cat > ./dist/$(BINARY_NAME)-macos-amd64-md5

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o ./dist/$(BINARY_NAME)-macos-arm64
	md5sum ./dist/$(BINARY_NAME)-macos-arm64 | head -c 32 | cat > ./dist/$(BINARY_NAME)-macos-arm64-md5

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(BINARY_NAME)-linux-amd64
	md5sum ./dist/$(BINARY_NAME)-linux-amd64 | head -c 32 | cat > ./dist/$(BINARY_NAME)-linux-amd64-md5

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o ./dist/$(BINARY_NAME)-linux-arm64
	md5sum ./dist/$(BINARY_NAME)-linux-arm64 | head -c 32 | cat > ./dist/$(BINARY_NAME)-linux-arm64-md5

clean:
	rm ./dist/${BINARY_NAME}-*
