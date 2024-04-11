run: build
	@./bin/goredis

build:
	@go build -o bin/goredis .
