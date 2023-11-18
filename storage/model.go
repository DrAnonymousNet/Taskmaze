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


func (t *Task) serialize() []byte {
	deadline := t.Deadline.Format(time.RFC3339)
	remindMe := t.RemindMe.Format(time.RFC3339)
	createdAt := t.CreatedAt.Format(time.RFC3339)

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

	createdAt, err := time.Parse(time.RFC3339, string(parts[1]))
	if err != nil {
		return fmt.Errorf("error parsing created at: %w", err)
	}
	t.CreatedAt = createdAt

	t.Title = string(parts[2])

	deadline, err := time.Parse(time.RFC3339, string(parts[3]))
	if err != nil {
		return fmt.Errorf("error parsing deadline: %w", err)
	}
	t.Deadline = deadline

	remindMe, err := time.Parse(time.RFC3339, string(parts[4]))
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
