run: build
	@./bin/uwe

build:
	@go build -o bin/uwe .

seed:
	@go run cmd/seed/main.go