all: notels

notels:
	go build .

generate:
	go generate -v ./...
