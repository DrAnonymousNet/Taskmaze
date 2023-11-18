package actions

import (
	"fmt"
	"time"

	"github.com/DrAnonymousNet/taskmaze/parser"
	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/validators"
)



func Create(title string, priority string, deadline string, remindMe string) (int, error) {
	var (
		deadlineTime time.Time
		remindMeTime time.Time
		err error
	)
	base := time.Now()


	if deadline != "" {
		deadlineTime, err = parser.Parse(deadline, base)
		if err != nil {
			return -1 , fmt.Errorf("error parsing deadline: %w", err)
		}
	}
	if remindMe != "" {
		remindMeTime, err = parser.Parse(remindMe, base, parser.WithDirection(parser.Future))
		if err != nil {
			return -1,  fmt.Errorf("error parsing remind me: %w", err)
		}
	}

	update := false
	if valid, err := validators.ValidateCreateArgs(title, deadlineTime, priority, remindMeTime, update); !valid {
		return -1, fmt.Errorf("error validating arguments: %w", err)
	}
	
	task := storage.CreatNewTask(title, deadlineTime, priority, remindMeTime)
	id, err := storage.AddTaskToDB(task)	
	return id, err
}

// For the due date, valid values include:
// - today
// - tomorrow
// - days of the week (e.g. monday, tuesday, wednesday, thursday, friday, saturday, sunday)
// In that case, the task should be due on the next day of the week
// - a date in the format YYYY-MM-DD
// 
