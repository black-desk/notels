/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/black-desk/notels/internal/utils/logger"
	"github.com/spf13/cobra"
)

var log = logger.Get("notels")

var rootCmd = &cobra.Command{
	Use:   "notels",
	Short: "a language server for markdown note taking",
	Long: `Notels is a language server for markdown note taking, compatible
with a subset syntax for a plugin enabled pandoc setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TODO\n")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
