.PHONY: build
build:
	CGO_ENABLED=0 go build -o elk-demo cmd/main.go

