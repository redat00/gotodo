package gotodo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCategoryCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("category create command")
	},
}

var deleteCategoryCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("category delete command")
	},
}

var getCategoryCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a category",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("category get command")
	},
}

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Interact with categories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command to interact with categories. Run with `--help` to get list of available command.")
	},
}

func init() {
	categoryCmd.AddCommand(addCategoryCmd, getCategoryCmd, deleteCategoryCmd)
	rootCmd.AddCommand(categoryCmd)
}
