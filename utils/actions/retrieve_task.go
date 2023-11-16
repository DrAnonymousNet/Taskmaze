package actions

import (
	"fmt"

	"github.com/DrAnonymousNet/taskmaze/storage"
	"github.com/DrAnonymousNet/taskmaze/utils/validators"
)

func Retrieve(id string) (*storage.Task, error) {
	if valid, err := validators.ValidateRetrieveArgs(id); !valid {
		fmt.Println(err)
		return nil, err
	}
	task, err := storage.RetrieveTaskFromDB(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return task, nil

}
