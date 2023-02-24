package gotodo

import (
	"github.com/redat00/gotodo/pkg/gotodo"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize gotodo",
	Run: func(cmd *cobra.Command, args []string) {
		gotodo.InitApp()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
