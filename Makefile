test:
	 go test ./...

go-imports:
	goimports -w .

build:  go-imports
	go build -o bin/ ./...