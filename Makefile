tidy:
	go mod tidy

run:
	go run cmd/server/main.go

build:
	cd cmd/server; go build -o ../../bin/go-gdrive

exec:
	./bin/go-gdrive

startapp: build exec