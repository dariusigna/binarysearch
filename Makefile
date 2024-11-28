.PHONY: test

test:
	go test ./...

start:
	go run cmd/api.go