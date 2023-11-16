package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
)


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Delete a task from the list of tasks",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//id := args[0]
		actions.Delete()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
