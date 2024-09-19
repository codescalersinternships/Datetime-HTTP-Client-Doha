format:
	gofmt -w .

build:
	go build -o ./client ./cmd/client.go

test:
	go test -v ./...

linter:
	golangci-lint run ./...