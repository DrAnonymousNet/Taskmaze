package actions

import (
	"fmt"
	"time"

	naturaldate "github.com/DrAnonymousNet/taskmaze/parser"
	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/utils"
	"github.com/DrAnonymousNet/taskmaze/validators"
)

func Update(id string, updateData map[string]interface{}) (int, error) {
	var (
		deadlineTime  time.Time
		remindMeTime  time.Time
		err           error
		remindMeExist bool
		deadlineExist bool
	)
	base := time.Now()

	deadline, deadlineExist := updateData[utils.DEADLINE]
	
	if deadlineExist {
		deadlineTime, err = naturaldate.Parse(deadline.(string), base)
		if err != nil {
			return -1, fmt.Errorf("error parsing deadline: %w", err)
		}
		updateData[utils.DEADLINE]=deadlineTime
	}

	remindMe, remindMeExist := updateData[utils.REMIND_ME]

	if remindMeExist {
		remindMeTime, err = naturaldate.Parse(remindMe.(string), base, naturaldate.WithDirection(naturaldate.Future))
		if err != nil {
			return -1, fmt.Errorf("error parsing remind me: %w", err)
		}
		updateData[utils.REMIND_ME] = remindMeTime
	}

	// since the deadline and remind me time depends on each other,
	// if the user wanted to update remindMe and does not intend to update
	// deadline, retrieve the deadline from the db so as to validate them together.
	// same with the deadline.
	task, err := storage.RetrieveTaskFromDB(id)
	if !remindMeExist {
		updateData[utils.REMIND_ME] = task.RemindMe
	}
	if !deadlineExist{
		updateData[utils.DEADLINE] = task.Deadline
	}
	if err != nil {
		return -1, fmt.Errorf("error updating task with id %s: %w", id, err)
	}

	if valid, err := validators.ValidateUpdateArgs(updateData); !valid {
		return -1, fmt.Errorf("error validating update arguments: %w", err)
	}

	return storage.UpdateTaskInDB(id, updateData)
}
