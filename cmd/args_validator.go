package cmd

import (
	"fmt"
	"time"
)

// IsTimeValid checks if a time string is a valid time.
func validateDueTime(dueTime string) bool {

	_, err := time.Parse("3:04PM", dueTime)
	if err == nil {
		return true
	}
	fmt.Printf("Error: %s\n", err)
	return false
}

func validateDueDate(dueDate string) bool {
	if dueDate == "today" {
		return true
	}
	if dueDate == "tomorrow" {
		return true
	}

	_, err := time.Parse("2006-01-02", dueDate)
	if err == nil {
		return true
	}
	fmt.Printf("Error: %s\n", err)
	return false
}

func validateReminder(reminder int, dueTime time.Time) bool {
	if reminder == 0 {
		return true
	}
	// The reminder + the current time should not be greatef that the dueTime
	if time.Now().Add(time.Duration(reminder) * time.Minute).After(dueTime) {
		return false
	}
	return true

}

func validatePriority(priority string) bool {
	if priority == "high" || priority == "medium" || priority == "low" {
		return true
	}
	return false
}


func ValidateCreateArgs(
	title string,
	dueDate string,
	priority string,
	dueTime string,
	reminder int,
) bool {
	var parsedDueDate time.Time
	if !validateDueDate(dueDate) {
		return false
	}
	if dueDate == "today" || dueDate == "tomorrow" {
		parsedDueDate = time.Now()
	}else{
		parsedDueDate, _ = time.Parse("2006-01-02", dueDate)
	}

	if !validateDueTime(dueTime) {
		return false
	}
	
	if !validateReminder(reminder, parsedDueDate) {
		return false

	}
	if !validatePriority(priority) {
		return false
	}
	return true
}
