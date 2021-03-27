test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

build-release:
	go build -ldflags "-s -w"
