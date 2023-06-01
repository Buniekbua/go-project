build:
	@go build -o bin/gousers

run: build
	@./bin/gousers

test:
	@go test -v ./...