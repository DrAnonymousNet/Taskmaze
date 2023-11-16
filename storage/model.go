package storage

import (
	"fmt"
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

	return []byte(fmt.Sprintf("%d %s %s %s %s %s %t", t.ID, createdAt, t.Title, deadline, remindMe, t.Priority, t.Done))

}

func (t *Task) deserialize(data []byte) error {
	var (
		deadline  string
		remindMe  string
		createdAt string
	)
	_, err := fmt.Sscanf(string(data), "%d %s %s %s %s %s %t", &t.ID, &createdAt, &t.Title, &deadline, &remindMe, &t.Priority, &t.Done)
	if err != nil {
		return fmt.Errorf("error deserializing task: %w", err)
	}
	t.Deadline, err = time.Parse("2006-01-02T15:04:05", deadline)
	if err != nil {
		return fmt.Errorf("error parsing deadline: %w", err)
	}
	t.RemindMe, err = time.Parse("2006-01-02T15:04:05", remindMe)
	if err != nil {
		return fmt.Errorf("error parsing remind me: %w", err)
	}
	t.CreatedAt, err = time.Parse("2006-01-02T15:04:05", createdAt)
	if err != nil {
		return fmt.Errorf("error parsing created at: %w", err)
	}

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
