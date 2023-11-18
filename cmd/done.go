package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/actions"
)


var DoneCMD = cobra.Command{
	Use: "done",
	Short: "mark task with given id as done",
	Example: "taskmaze done 23",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := actions.CompleteTask(id)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Task with id %s marked as completed successfully\n", id)
	},
}

func init(){
	RootCmd.AddCommand(&DoneCMD)

}