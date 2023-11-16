package tests

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"
//
//
//
//

// 	"github.com/DrAnonymousNet/taskmaze/cmd"
// 	"github.com/DrAnonymousNet/taskmaze/storage"
// )

// func TestCreateTaskCmd(t *testing.T) {
// 	// Create a temporary file to use as the database
// 	tmpfile, err := ioutil.TempFile("", "testdb")
// 	if err != nil {
// 		t.Fatalf("error creating temporary file: %v", err)
// 	}
// 	defer os.Remove(tmpfile.Name())

// 	// Initialize the database
// 	storage.InitDB(tmpfile.Name())

// 	// Initialize the actions package with the database

// 	// Create a new task command
// 	cmd := cmd.CreateTaskCmd()

// 	// Test creating a new task with valid arguments
// 	args := []string{"--title", "Test Task", "--priority", "high", "--deadline", "tomorrow by 3pm", "--remind_me", "tomorrow by 2pm"}
// 	cmd.SetArgs(args)
// 	err = cmd.Execute()
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}

// 	// Test creating a new task with invalid arguments
// 	args = []string{"--title", "Test Task", "--priority", "invalid", "--due_date", "2022-01-01", "--reminder", "1h"}
// 	cmd.SetArgs(args)
// 	err = cmd.Execute()
// 	if err == nil {
// 		t.Errorf("expected error, but got nil")
// 	}
// }
