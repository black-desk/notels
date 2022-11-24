.PHONY: cloc check-license

all:
	go build .

cloc:
	cloc . --not-match-f='_gen.go$$' --exclude-ext='md,json'

check-license:
	go run github.com/google/go-licenses@latest report .
	go run github.com/google/go-licenses@latest check .
