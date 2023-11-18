package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/utils"
	"github.com/DrAnonymousNet/taskmaze/utils/actions"
)


var (
	listTaskCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long:  "List all tasks that satisfies the given criteria",
		Example: `taskMaze list [--done|--not-done|--missed-deadline|--due-today|--due-tomorrow|--due-this-week|--due-this-month]`,
		Run: func(cmd *cobra.Command, args []string) {
			filters := make(map[string]interface{})
			if doneFlag {
				filters[utils.DONE] = true
			}
			if notDoneFlag {
				filters[utils.NOT_DONE] = true
			}
			if missedDeadlineFlag {
				filters[utils.MISSED_DEADLINE] = true
			}
			if dueTodayFlag {
				filters[utils.DUE_TODAY] = true
			}
			if dueTomorrowFlag {
				filters[utils.DUE_TOMORROW] = true
			}
			if dueThisWeekFlag {
				filters[utils.DUE_THIS_WEEK] = true
			}
			if dueThisMonthFlag {
				filters[utils.DUE_THIS_MONTH] = true
			}
			

			tasks, err := actions.List(filters)
			if err!= nil {
				fmt.Println(err.Error())
				return
			}
			storage.DisplayManyTasks(tasks)
			// for _, task := range tasks {
			// 	task.Display()
			// }
		},
	}

	doneFlag         bool
	notDoneFlag      bool
	missedDeadlineFlag   bool
	dueTodayFlag     bool
	dueTomorrowFlag  bool
	dueThisWeekFlag  bool
	dueThisMonthFlag bool
	orderFlag        string
)

func init() {
	listTaskCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "List tasks that are done")
	listTaskCmd.Flags().BoolVarP(&notDoneFlag, "not-done", "n", false, "List tasks that are not done")
	listTaskCmd.Flags().BoolVarP(&missedDeadlineFlag, "missed-deadline", "m", false, "List tasks with missed deadlines")
	listTaskCmd.Flags().BoolVarP(&dueTodayFlag, "due-today", "t", false, "List tasks due today")
	listTaskCmd.Flags().BoolVarP(&dueTomorrowFlag, "due-tomorrow", "", false, "List tasks due tomorrow")
	listTaskCmd.Flags().BoolVarP(&dueThisWeekFlag, "due-this-week", "", false, "List tasks due this week")
	listTaskCmd.Flags().BoolVarP(&dueThisMonthFlag, "due-this-month", "", false, "List tasks due this month")
	listTaskCmd.Flags().StringVarP(&orderFlag, "order", "r", "", "Order tasks by priority, deadline, created, updated")

	RootCmd.AddCommand(listTaskCmd)
}
