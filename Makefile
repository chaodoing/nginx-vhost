BINARY_NAME=vhost
DATE=`date +%FT%T%z`
VERSION=1.0.0

build:
	@echo "build apple-darwin version: ${VERSION} date: ${DATE}"
	@GOOS=darwin go build -o release/apple_darwin/$(BINARY_NAME) ./
windows:
	@echo "build windows-i386 version: ${VERSION} date: ${DATE}"
	@GOOS=windows GOARCH=386 go build -o  release/windows/i386/$(BINARY_NAME) ./
	@echo "build windows-amd64 version: ${VERSION} date: ${DATE}"
	@GOOS=windows GOARCH=amd64 go build -o  release/windows/amd64/$(BINARY_NAME) ./
linux:
	@echo "build linux-amd64 version: ${VERSION} date: ${DATE}"
	@GOOS=linux GOARCH=amd64 go build -o  release/linux_amd64/$(BINARY_NAME) ./
	@echo "build linux-i386 version: ${VERSION} date: ${DATE}"
	@GOOS=linux GOARCH=386 go build -o  release/linux/i386/$(BINARY_NAME) ./
all:
	@echo "build apple-darwin version: ${VERSION} date: ${DATE}"
	@GOOS=darwin go build -o release/apple_darwin/$(BINARY_NAME) ./
	@echo "build linux-amd64 version: ${VERSION} date: ${DATE}"
	@GOOS=linux GOARCH=amd64 go build -o  release/linux/amd64/$(BINARY_NAME) ./
	@echo "build linux-i386 version: ${VERSION} date: ${DATE}"
	@GOOS=linux GOARCH=386 go build -o  release/linux/i386/$(BINARY_NAME) ./
	@echo "build windows-i386 version: ${VERSION} date: ${DATE}"
	@GOOS=windows GOARCH=386 go build -o  release/windows/i386/$(BINARY_NAME) ./
	@echo "build windows-amd64 version: ${VERSION} date: ${DATE}"
	@GOOS=windows GOARCH=amd64 go build -o  release/windows/amd64/$(BINARY_NAME) ./
clean:
	@rm -rf release
	@rm -rf arrange