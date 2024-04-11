run: build
	@./bin/goredis --listenAddr :5001

build:
	@go build -o bin/goredis .
