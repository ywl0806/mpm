.PHONY: build run

build: 
	go build -o mpm && mv ./mpm ~/go/bin

dev:
	ENV=dev go run . $(ARGS) 