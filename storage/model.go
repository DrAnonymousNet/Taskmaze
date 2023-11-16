package storage

import (
	"errors"
	"time"
)


type Task struct {
	CreatedAt time.Time
	Title string
	DueDate string
	Priority string
	Time string
	Reminder string
	Done bool
}

func CreatNewTask(title string, dueDate string, priority string, time string, reminder string) Task {
	task := Task{
		Title: title,
		DueDate: dueDate,
		Priority: priority,
		Time: time,
		Reminder: reminder,
		Done: false,
	}
	return task
}

func AddTaskToDB(task Task) int {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	nextID := len(TaskDB.GlobalTasksMap)
	TaskDB.GlobalTasksMap[nextID] = task
	return nextID
}

func GetTaskByID(id int) (Task, error) {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	task, ok := TaskDB.GlobalTasksMap[id]
	if!ok {
		return Task{}, errors.New("Task not found")
	}
	return task, nil
}

func DeleteTask(id int) {
	TaskDB.mu.Lock()
	defer TaskDB.mu.Unlock()
	delete(TaskDB.GlobalTasksMap, id)
}

	