package gotodo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task create command")
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task delete command")
	},
}

var getTaskCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task get command")
	},
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Interact with tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command to interact with tasks. Run with `--help` to list available commands.")
	},
}

func init() {
	taskCmd.AddCommand(addTaskCmd, deleteTaskCmd, getTaskCmd)
	rootCmd.AddCommand(taskCmd)
}
