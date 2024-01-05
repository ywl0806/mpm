.PHONY: build run

build: build
	go build -o mpm

run:
	ENV=dev go run . $(ARGS)