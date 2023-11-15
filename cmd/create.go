package cmd


import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var title string
var dueDate string
var priority string
var deadline string
var remind_me_by int

var createTaskCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new task",
	Long: "Create a new task to be done",
	Run: func (cmd *cobra.Command, args []string)  {
		
	},
	ValidArgs: []string{"title", "due_date", "priority", "time", "reminder"},
}

func init() {
	rootCmd.AddCommand(createTaskCmd)
	createTaskCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	createTaskCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority of the task")
	createTaskCmd.Flags().StringVarP(&deadline, "time", "T", "", "deadline of the task")
	createTaskCmd.Flags().IntVarP(&remind_me_by, "remind_me_by", "r", 0, "When you should be remind of the task")
	createTaskCmd.MarkFlagRequired("title")

}
