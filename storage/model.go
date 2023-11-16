package storage

import (
	"errors"
	"fmt"
	"time"
)


type Task struct {
	ID int
	CreatedAt time.Time
	Title string
	Deadline time.Time
	Priority string
	RemindMe time.Time
	Done bool
}

func CreatNewTask(title string, deadline time.Time, priority string, remindMe time.Time) Task {
	task := Task{

		CreatedAt: time.Now(),
		Title: title,
		Deadline: deadline,
		Priority: priority,
		RemindMe: remindMe,
		Done: false,
	}
	return task
}

func AddTaskToDB(task *Task) int {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	nextID := len(*TaskDB.GlobalTasksMap)
	fmt.Println(nextID, "Next ID")
	(*TaskDB.GlobalTasksMap)[nextID] = *task
	return nextID
}

func GetTaskByID(id int) (Task, error) {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	task, ok := (*TaskDB.GlobalTasksMap)[id]
	if!ok {
		return Task{}, errors.New("Task not found")
	}
	return task, nil
}

func DeleteTask(id int) {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	delete(*TaskDB.GlobalTasksMap, id)
}

	