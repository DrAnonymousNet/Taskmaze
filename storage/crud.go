package storage

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"

	"github.com/DrAnonymousNet/taskmaze/utils"
)

func CreatNewTask(title string, deadline time.Time, priority string, remindMe time.Time) *Task {
	task := Task{

		CreatedAt: time.Now(),
		Title:     title,
		Deadline:  deadline,
		Priority:  priority,
		RemindMe:  remindMe,
		Done:      false,
	}
	return &task
}

func UpdateTask(task *Task, data map[string]interface{}) {
	for field, value := range data {
		switch field {
		case utils.TITLE:
			task.Title = value.(string)
		case utils.DEADLINE:
			task.Deadline = value.(time.Time)
		case utils.PRIORITY:
			task.Priority = value.(string)
		case utils.REMIND_ME:
			task.RemindMe = value.(time.Time)
		case utils.DONE_FIELD:
			task.Done = value.(bool)
		}
	}
}

func AddTaskToDB(task *Task) (int, error) {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	err := MyDB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		task.ID = int(id)
		idBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(idBytes, uint64(id))
		b.Put(idBytes, task.serialize())
		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("error adding task to db: %w", err)
	}
	return task.ID, nil
}

func UpdateTaskInDB(task *Task) (int, error) {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	err := MyDB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id := task.ID
		idBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(idBytes, uint64(id))
		b.Put(idBytes, task.serialize())
		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("error updating task in db: %w", err)
	}
	return task.ID, nil
}

func RetrieveTaskFromDB(id string) (*Task, error) {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	var task Task
	err := MyDB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		idBytes := make([]byte, 8)
		id, _ := strconv.Atoi(id) // Already validated
		binary.BigEndian.PutUint64(idBytes, uint64(id))
		data := b.Get(idBytes)
		if data == nil {
			return fmt.Errorf("task not found")
		}
		err := task.deserialize(data)
		if err != nil {
			return fmt.Errorf("error deserializing task: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error retrieving task from db: %w", err)
	}
	return &task, nil
}

func DeleteTaskFromDB(id string) error {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	err := MyDB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		idBytes := make([]byte, 8)
		id, _ := strconv.Atoi(id) // Already validated
		binary.BigEndian.PutUint64(idBytes, uint64(id))
		data := b.Get(idBytes)
		if data == nil {
			return fmt.Errorf("task not found")
		}
		err := b.Delete(idBytes)
		if err != nil {
			return fmt.Errorf("error deleting task: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting task from db: %w", err)
	}
	return nil
}

func ListTasksFromDB(query map[string]interface{}) ([]*Task, error) {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	var tasks []*Task
	err := MyDB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := task.deserialize(v)
			if err != nil {
				return fmt.Errorf("error deserializing task: %w", err)
			}
			shouldReturn := applyFilters(&task, query)
			if shouldReturn {
				tasks = append(tasks, &task)
			}
		}
		return nil
	})

	if err != nil {
		return []*Task{}, fmt.Errorf("error listing tasks from db: %w", err)
	}
	return tasks, nil
}

func MarkTaskAsComplete(id string) error {
	task, err := RetrieveTaskFromDB(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve task in db: %w", err)
	}
	UpdateTask(task, map[string]interface{}{utils.DONE_FIELD: true})
	_, err = UpdateTaskInDB(task)
	if err != nil{
		return fmt.Errorf("failed to update task as complete in db: %w", err)
	}
	return nil


}
