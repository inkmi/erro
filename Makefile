build:  go-imports
	go build -o bin/ ./...

test:
	 go test ./...

go-imports:
	goimports -w .

upgrade-deps:
	go get -u ./...
	go mod tidy
	gotestsum ./...

lint: staticcheck
	golangci-lint run

audit:
	go list -json -deps ./... | nancy sleuth --loud

sec: audit
	gosec  .
	govulncheck ./...

staticcheck:
	staticcheck  .
