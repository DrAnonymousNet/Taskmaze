package cmd

import (
	"github.com/DrAnonymousNet/taskmaze/utils/actions"
	"github.com/spf13/cobra"
)

var (
	listTaskCmd = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long:  "List all tasks that satisfies the given criteria",
		Run: func(cmd *cobra.Command, args []string) {
			actions.List()
		},
	}

	doneFlag         bool
	notDoneFlag      bool
	overdueFlag      bool
	missedDeadline   bool
	dueTodayFlag     bool
	dueTomorrowFlag  bool
	dueThisWeekFlag  bool
	dueThisMonthFlag bool
	orderFlag        string
)

func init() {
	listTaskCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "List tasks that are done")
	listTaskCmd.Flags().BoolVarP(&notDoneFlag, "not-done", "n", false, "List tasks that are not done")
	listTaskCmd.Flags().BoolVarP(&overdueFlag, "overdue", "o", false, "List tasks with deadlines greater than now")
	listTaskCmd.Flags().BoolVarP(&missedDeadline, "missed-deadline", "m", false, "List tasks with missed deadlines")
	listTaskCmd.Flags().BoolVarP(&dueTodayFlag, "due-today", "t", false, "List tasks due today")
	listTaskCmd.Flags().BoolVarP(&dueTomorrowFlag, "due-tomorrow", "", false, "List tasks due tomorrow")
	listTaskCmd.Flags().BoolVarP(&dueThisWeekFlag, "due-this-week", "", false, "List tasks due this week")
	listTaskCmd.Flags().BoolVarP(&dueThisMonthFlag, "due-this-month", "", false, "List tasks due this month")
	listTaskCmd.Flags().StringVarP(&orderFlag, "order", "r", "", "Order tasks by priority, deadline, created, updated")
	listTaskCmd.MarkFlagsMutuallyExclusive([]string{"done", "not-done"}...)
	RootCmd.AddCommand(listTaskCmd)
}
