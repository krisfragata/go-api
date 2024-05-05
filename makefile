build: 
	@go build -o bin/go-api-flow

run: build 
	@./bin/go-api-flow

test:
	@go  test -v ./...