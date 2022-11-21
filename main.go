/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/black-desk/notels/cmd"

func main() {
	cmd.Execute()
	X := []string{}
	if err := func() error {
		for _, x := range X {
			if "" == x {
				return nil
			}
		}

		return nil
	}(); err != nil {
		return err
	}

}
