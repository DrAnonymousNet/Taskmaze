package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
)

// createCmd represents the create command
var title string
var priority string
var deadline string
var remind_me string

var createTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  "Create a new task to be done",
	Example: `taskMaze create --title "Buy milk" --priority "high" --deadline "today by 3pm" --remind_me "today by 1pm"`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := actions.Create(title, priority, deadline, remind_me)
		if err != nil {
			fmt.Println(fmt.Errorf("error creating task: %v", err).Error())
		} else {
			fmt.Println("Task created with id: ", id)
		}
	},
	ValidArgs: []string{"title", "due_date", "priority", "time", "reminder"},
}

func init() {
	RootCmd.AddCommand(createTaskCmd)
	createTaskCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	createTaskCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority of the task")
	createTaskCmd.Flags().StringVarP(&deadline, "deadline", "T", "", "deadline of the task")
	createTaskCmd.Flags().StringVarP(&remind_me, "remind_me", "r", "", "When you should be remind of the task")
	createTaskCmd.MarkFlagRequired("title")

}
