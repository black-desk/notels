package cmd

import (
	"os"
	"path/filepath"

	"github.com/black-desk/notels/internal/export"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [path/to/workspace [path/to/output]]",
	Short: "render markdown notes in workspace to html",
	Run: func(cmd *cobra.Command, args []string) {

		workspacePath := ""
		outputPath := ""
		var err error

		switch len(args) {
		case 2:
			outputPath = args[1]
			fallthrough
		case 1:
			workspacePath = args[0]
		case 0:
		default:
			log.Fatalw("too many arguments")
		}

		if workspacePath == "" {
			workspacePath, err = os.Getwd()
			if err != nil {
				log.Fatalw("failed to get working dir",
					"error", err)
			}
		}

		if outputPath == "" {
			outputPath = filepath.Join(workspacePath,
				"__OUTPUT__")
		}

		err = export.Export(workspacePath, outputPath)
		if err != nil {
			log.Fatalw("failed to export workspace",
				"error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
