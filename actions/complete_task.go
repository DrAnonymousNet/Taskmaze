package actions

import (
	"fmt"

	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/validators"
)

func CompleteTask(id string) error {
	if valid, err := validators.ValidateCompleteArgs(id); !valid || err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}
	err := storage.MarkTaskAsComplete(id)
	if err != nil{
		return fmt.Errorf("failed to complete task: %w", err)
	}
	return nil
}
