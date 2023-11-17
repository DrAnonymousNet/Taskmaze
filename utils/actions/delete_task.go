package actions

import (
	"fmt"

	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/utils/validators"
)


func Delete(id string) error {
	if valid, err := validators.ValidateDeleteArgs(id);!valid {
		return fmt.Errorf("failed to delete task %w", err)
	}
	if err := storage.DeleteTaskFromDB(id); err!= nil {
		return fmt.Errorf("failed to delete task %w", err)
	}
	return nil
}