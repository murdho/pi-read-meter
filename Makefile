test:
	@go test -v

build:
	@go build

run: build
	@./pi-read-meter config.dev.json
