OUTPUT_FILE := version.txt

GET_TAG_CMD := git tag --sort=committerdate | grep -E '[0-9]' | tail -1 | cut -b 2-7

.PHONY: write_tag

write_tag:
	@echo $(shell $(GET_TAG_CMD)) > $(OUTPUT_FILE)

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
