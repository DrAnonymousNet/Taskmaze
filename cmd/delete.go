package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/utils/actions"
)

var (
	showTaskAndConfirmDeleteFlag bool
	
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Delete a task from the list of tasks",
	Example: `taskMaze delete 1 [--show]`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		if showTaskAndConfirmDeleteFlag {
			task, err := actions.Retrieve(id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			task.Display()
			fmt.Print("Are you sure you want to delete this task? [y/n]: ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm!= "y" {
				fmt.Println("Task not deleted")
				return
			}
		}

		if err := actions.Delete(id); err!= nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Task %s deleted successfully\n", id)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&showTaskAndConfirmDeleteFlag, "show", "s", false, "Show the task before deleting it and prompt confirmation")
	RootCmd.AddCommand(deleteCmd)
}
