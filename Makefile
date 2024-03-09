lint:
	golangci-lint run

test:
	go test -v -cover ./... -coverprofile=coverage.out

coverage: test
	go tool cover -html=coverage.out
