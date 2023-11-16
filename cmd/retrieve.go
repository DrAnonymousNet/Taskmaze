package cmd

import (
	"github.com/DrAnonymousNet/taskmaze/utils/actions"
	"github.com/spf13/cobra"
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
	RootCmd.AddCommand(retrieveTaskCmd)
}
