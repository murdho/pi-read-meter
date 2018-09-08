test:
	@go test .

build:
	@go build

run: build
	@./pi-read-meter
