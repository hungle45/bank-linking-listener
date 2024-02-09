run:
	go run ./cmd/app

lint:
	golangci-lint run ./...

test:
	go test -v -cover -race ./...

