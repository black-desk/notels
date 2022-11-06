package cmd

import (
	"github.com/black-desk/notels/internal/watch"
	"github.com/spf13/cobra"
	"os"
)

var watchCmdDaemon bool

var watchCmd = &cobra.Command{
	Use:   "watch [path/to/workspace]",
	Short: "watch markdown notes in workspace, and render them to a html then service them",
	Run: func(cmd *cobra.Command, args []string) {

		workspacePath := ""
		var err error

		switch len(args) {
		case 1:
			workspacePath = args[0]
		case 0:
		default:
			log.Fatalw("too many arguments")
		}

		if watchCmdDaemon {
			log.Fatalw("TODO")
		}

		if workspacePath == "" {
			workspacePath, err = os.Getwd()
			if err != nil {
				log.Fatalw("failed to get working dir",
					"error", err)
			}
		}

		err = watch.Watch(workspacePath)
		if err != nil {
			log.Fatalw("error occurred during watching workspace",
				"error", err)
		}
	},
}

func init() {
	watchCmd.Flags().BoolVarP(&watchCmdDaemon,
		"daemon", "d", false, "daemonlize or not")

	rootCmd.AddCommand(watchCmd)
}
