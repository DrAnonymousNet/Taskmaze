package storage

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"
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

func (t *Task) Display() {
	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Printf("║  Task Details:                    ║\n")
	fmt.Println("╠═══════════════════════════════════╣")
	fmt.Printf("║  ID:        %d\n", t.ID)
	fmt.Printf("║  Title:     %s\n", t.Title)
	fmt.Printf("║  Due Date:  %s\n", t.Deadline.Format("2006-01-02 by 15:04"))
	fmt.Printf("║  Priority:  %s\n", t.Priority)
	fmt.Printf("║  Remind Me: %s\n", t.RemindMe.Format("2006-01-02 by 15:04"))
	fmt.Printf("║  Completed: %t\n", t.Done)
	fmt.Println("╚═══════════════════════════════════╝")

}

func (t *Task) serialize() []byte {
	deadline := t.Deadline.Format("2006-01-02T15:04:05")
	remindMe := t.RemindMe.Format("2006-01-02T15:04:05")
	createdAt := t.CreatedAt.Format("2006-01-02T15:04:05")

	return []byte(fmt.Sprintf("%d|%s|%s|%s|%s|%s|%t", t.ID, createdAt, t.Title, deadline, remindMe, t.Priority, t.Done))

}
func (t *Task) deserialize(data []byte) error {
	parts := bytes.Split(data, []byte{'|'})
	if len(parts) != 7 {
		return errors.New("invalid task data")
	}

	id, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return fmt.Errorf("error parsing ID: %w", err)
	}
	t.ID = id

	createdAt, err := time.Parse("2006-01-02T15:04:05", string(parts[1]))
	if err != nil {
		return fmt.Errorf("error parsing created at: %w", err)
	}
	t.CreatedAt = createdAt

	t.Title = string(parts[2])

	deadline, err := time.Parse("2006-01-02T15:04:05", string(parts[3]))
	if err != nil {
		return fmt.Errorf("error parsing deadline: %w", err)
	}
	t.Deadline = deadline

	remindMe, err := time.Parse("2006-01-02T15:04:05", string(parts[4]))
	if err != nil {
		return fmt.Errorf("error parsing remind me: %w", err)
	}
	t.RemindMe = remindMe

	t.Priority = string(parts[5])

	done, err := strconv.ParseBool(string(parts[6]))
	if err != nil {
		return fmt.Errorf("error parsing done: %w", err)
	}
	t.Done = done

	return nil
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
