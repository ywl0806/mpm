.PHONY: build run

VERSION=1.0.0

# build: 
# 	go build -o mpm && mv ./mpm ~/go/bin

dev:
	ENV=dev go run . $(ARGS) 


build-release:
	GOOS=darwin GOARCH=amd64 go build -o mpm
	tar -czvf mpm-darwin-amd64.tar.gz mpm
	rm mpm

	GOOS=darwin GOARCH=arm64 go build -o mpm
	tar -czvf mpm-darwin-arm64.tar.gz mpm
	rm mpm

	mkdir -p ./releases/download/${VERSION}
	mv mpm-darwin-amd64.tar.gz ./releases/download/${VERSION}/mpm-darwin-amd64.tar.gz
	mv mpm-darwin-arm64.tar.gz ./releases/download/${VERSION}/mpm-darwin-arm64.tar.gz

	shasum -a 256 ./releases/download/${VERSION}/mpm-darwin-amd64.tar.gz > ./releases/download/${VERSION}/mpm-darwin-amd64.tar.gz.sha256
	@echo "SHA256 for mpm-darwin-amd64.tar.gz"
	@echo "\n\n-------------------------------------\n\n"
	cat ./releases/download/${VERSION}/mpm-darwin-amd64.tar.gz.sha256
	@echo "\n\n-------------------------------------\n\n"
	rm ./releases/download/${VERSION}/mpm-darwin-amd64.tar.gz.sha256

	shasum -a 256 ./releases/download/${VERSION}/mpm-darwin-arm64.tar.gz > ./releases/download/${VERSION}/mpm-darwin-arm64.tar.gz.sha256
	@echo "SHA256 for mpm-darwin-arm64.tar.gz"
	@echo "\n\n-------------------------------------\n\n"
	cat ./releases/download/${VERSION}/mpm-darwin-arm64.tar.gz.sha256
	@echo "\n\n-------------------------------------\n\n"
	rm ./releases/download/${VERSION}/mpm-darwin-arm64.tar.gz.sha256
