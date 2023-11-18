package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
	"github.com/DrAnonymousNet/taskmaze/utils"
)

var done bool

var UpdateTaskCmd = &cobra.Command{
	Use:     "update",
	Short:   "update an existing task",
	Long:    "update an task to be done",
	Example: `taskMaze update --title "Buy milk" --priority "high" --deadline "today by 3pm" --remind_me "today by 1pm"`,
	Run: func(cmd *cobra.Command, args []string) {
		updateData := make(map[string]interface{})
		if title != ""{
			updateData[utils.TITLE] = title
		}
		if priority != ""{
			updateData[utils.PRIORITY] = priority
		}
		if deadline != ""{
			updateData[utils.DEADLINE] = deadline
 		}
		if remind_me != ""{
			updateData[utils.REMIND_ME] = remind_me 
		}
		if done{
			updateData[utils.DONE_FIELD] = done
		}
		id := args[0]
		taskId, err := actions.Update(id, updateData)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Task with id %d updated successfully\n", taskId)
	},

}
func init(){
	RootCmd.AddCommand(UpdateTaskCmd)
	UpdateTaskCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	UpdateTaskCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority of the task")
	UpdateTaskCmd.Flags().StringVarP(&deadline, "deadline", "T", "", "deadline of the task")
	UpdateTaskCmd.Flags().StringVarP(&remind_me, "remind-me", "r", "", "When you should be remind of the task")
	UpdateTaskCmd.Flags().BoolVarP(&done, "done", "d", false, "mark the task as done")
}
