package main

import "github.com/black-desk/notels/cmd"

func main() {
	cmd.Execute()
}

//go:generate go run github.com/ribice/glice/v2/cmd/glice@latest -f -i
