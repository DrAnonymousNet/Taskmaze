package storage

import (
	"fmt"
	"time"

	"github.com/DrAnonymousNet/taskmaze/utils"
)

func applyNotDoneFilter(task *Task, flagValue bool) bool {
	if flagValue {
		return !task.Done
	}
	return true
}

func applyDoneFilter(task *Task, flagValue bool) bool {
	if flagValue {
		return task.Done
	}
	return true
}

func applyOverdueTodayFilter(task *Task, flagValue bool) bool {
	now := time.Now()
	if flagValue {
		return task.Deadline.After(
			time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				0, 0, 0, 0,
				now.Location(),
			)) && task.Deadline.Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Now().Location()))
	}
	return true
}

func applyOverdueTomorrowFilter(task *Task, flagValue bool) bool {
	if flagValue {
		return task.Deadline.After(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Now().Location())) && task.Deadline.Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, 0, 0, 0, 0, time.Now().Location()))
	}
	return true
}

func applyMissedDeadLineFilter(task *Task, flagValue bool) bool {
	if flagValue {
		return task.Deadline.Before(time.Now()) && !task.Done
	}
	return true
}
func applyDueThisWeekFilter(task *Task, flagValue bool) bool {
	if flagValue {
		currYear, currWeek := time.Now().ISOWeek()
		deadlineYear, deadlineWeek := task.Deadline.ISOWeek()
		return currYear == deadlineYear && deadlineWeek == currWeek
	}
	return true
}

func applyDueThisMonthFilter(task *Task, flagValue bool) bool {
	if flagValue {
		currMonth := time.Now().Month()
		currYear := time.Now().Year()
		deadlineMonth := task.Deadline.Month()
		deadlineYear := task.Deadline.Year()
		return currMonth == deadlineMonth && deadlineYear == currYear
	}
	return true
}

func applyFilters(task *Task, filters map[string]interface{}) bool {
	for flag, flagValue := range filters {
		switch flag {
		case utils.NOT_DONE:
			fmt.Println("applying not done filter")
			if !applyNotDoneFilter(task, flagValue.(bool)) {
				return false
			}
		case utils.DONE:
			fmt.Println("applying done filter")
			if !applyDoneFilter(task, flagValue.(bool)) {
				return false
			}
		case utils.DUE_TODAY:
			fmt.Println("applying due today filter")
			if !applyOverdueTodayFilter(task, flagValue.(bool)) {
				return false
			}
		case utils.DUE_TOMORROW:
			fmt.Println("applying due tomorrow filter")
			if !applyOverdueTomorrowFilter(task, flagValue.(bool)) {
				return false
			}
		case utils.MISSED_DEADLINE:
			fmt.Println("applying missed deadline filter")
			if !applyMissedDeadLineFilter(task, flagValue.(bool)) {
				fmt.Println("returning false")
				return false
			}
		case utils.DUE_THIS_WEEK:
			if !applyDueThisWeekFilter(task, flagValue.(bool)) {
				return false
			}
		case utils.DUE_THIS_MONTH:
			if !applyDueThisMonthFilter(task, flagValue.(bool)){
				return false
			}
		}
	}

	return true
}
