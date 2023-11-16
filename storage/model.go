package storage

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	ID        int
	CreatedAt time.Time
	Title     string
	Deadline  time.Time
	Priority  string
	RemindMe  time.Time
	Done      bool
}

func (t *Task) serialize() []byte {
	return []byte(fmt.Sprintf("%d|%s|%s|%s|%s|%t", t.ID, t.CreatedAt, t.Title, t.Deadline, t.Priority, t.Done))
}

func (t *Task) deserialize(data []byte) error {
	_, err := fmt.Sscanf(string(data), "%d|%s|%s|%s|%s|%t", &t.ID, &t.CreatedAt, &t.Title, &t.Deadline, &t.Priority, &t.Done)
	if err != nil {
		return err
	}
	return nil
}

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
		b.Put(idBytes, task.serialize())  
		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("error adding task to db: %w", err)
	}
	return task.ID, nil
}



// func GetTaskByID(id int) (Task, error) {
// 	TaskDB.mu.Lock()
// 	defer TaskDB.mu.Unlock()
// 	task, ok := (*TaskDB.TaskDB)[id]
// 	if !ok {
// 		return Task{}, errors.New("Task not found")
// 	}
// 	return task, nil
// }

// func DeleteTask(id int) {
// 	TaskDB.mu.Lock()
// 	defer TaskDB.mu.Unlock()
// 	delete(*TaskDB.TaskDB, id)
// }
