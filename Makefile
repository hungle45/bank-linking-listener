.PHONY: run-% lint test

run-%:
	go run ./cmd/$*

lint:
	golangci-lint run ./...

test:
	go test -v -cover -race ./...

