package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
)

var retrieveTaskCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a task",
	Long:  "Retrieve a task from the list of tasks",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//id := args[0]
		actions.Retrieve()
	},
}

func init() {
	rootCmd.AddCommand(retrieveTaskCmd)
}