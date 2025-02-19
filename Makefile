build:
	@go build -o bin/go-channels

run: build
	@./bin/go-channels

test:
	@go test -v ./...