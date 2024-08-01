build:
	@go build -o bin/go-bank
run: build
	@./bin/go-bank
test:
	@go test -v ./...
dep:
	@go mod vendor
	@go mod tidy
	@go mod verify