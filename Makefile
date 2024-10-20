.SILENT:

run:
	go build -o ./bin/main-test ./cmd/core/main.go
	./bin/main.go

test:
	go test -v ./...

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/main-linux-amd64 ./cmd/core/main.go
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/main-freebsd-amd64 ./cmd/core/main.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/main-windows-amd64.exe ./cmd/core/main.go