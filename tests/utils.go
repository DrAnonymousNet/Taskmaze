package tests

import (
	"fmt"
	"os"

	"github.com/DrAnonymousNet/taskmaze/storage"
)

// TeardownTestDB is a helper function to delete the test database
func TeardownTestDB() {
	storage.MyDB.DB.Close()
	err := os.Remove("test.db")
	if err != nil {
		fmt.Println(err)
	}
}
