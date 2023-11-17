package storage

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
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

func AddTaskToDB(task *Task) (int, error) {
	MyDB.mu.Lock()
	defer MyDB.mu.Unlock()
	err := MyDB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		task.ID = int(id)
		idBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(idBytes, uint64(id))
		fmt.Println(idBytes)
		b.Put(idBytes, task.serialize())  
		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("error adding task to db: %w", err)
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
	fmt.Println(task)
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