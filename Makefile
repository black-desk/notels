.PHONY: cloc

all: notels

notels:
	go build .

cloc:
	cloc . --not-match-f='_gen.go$$' --exclude-ext='md,json'

generate:
	go generate -v ./...
