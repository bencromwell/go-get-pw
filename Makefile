test:
	go test ./...
build-release:
	go build -ldflags "-s -w"
