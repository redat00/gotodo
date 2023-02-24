package gotodo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addStatusCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status create command")
	},
}

var deleteStatusCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status delete command")
	},
}

var getStatusCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status get command")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Interact with statuses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status command")
	},
}

func init() {
	statusCmd.AddCommand(getStatusCmd, deleteStatusCmd, getStatusCmd)
	rootCmd.AddCommand(statusCmd)
}
