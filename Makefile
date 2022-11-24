.PHONY: generate
all: notels CLOC

notels:
	go build .

CLOC:
	cloc . --not-match-f='_gen.go$$' --exclude-ext='md,json' > CLOC

generate:
	go generate ./...
