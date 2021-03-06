test:
	@go test -v

build:
	@go build

build-pi:
	@GOOS=linux GOARCH=arm GOARM=6 go build -o pi-read-meter-armv6

run: build
	@./pi-read-meter config.dev.json

clean:
	@rm -rf tmp
	@rm -f ./pi-read-meter
	@rm -f ./pi-read-meter-armv6
