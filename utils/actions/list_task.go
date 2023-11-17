package actions

import "github.com/DrAnonymousNet/taskmaze/storage"


func List(filters map[string]interface{}) ( []*storage.Task ,error) {
		
	return storage.ListTasksFromDB(filters)
}