package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
)

// createCmd represents the create command
var title string
var priority string
var deadline string
var remind_me_by string

var createTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  "Create a new task to be done",
	Run: func(cmd *cobra.Command, args []string) {
		actions.Create(title, priority, deadline, remind_me_by)
	},
	ValidArgs: []string{"title", "due_date", "priority", "time", "reminder"},
}

func init() {
	rootCmd.AddCommand(createTaskCmd)
	createTaskCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	createTaskCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority of the task")
	createTaskCmd.Flags().StringVarP(&deadline, "time", "T", "", "deadline of the task")
	createTaskCmd.Flags().StringVarP(&remind_me_by, "remind_me_by", "r", "", "When you should be remind of the task")
	createTaskCmd.MarkFlagRequired("title")

}
