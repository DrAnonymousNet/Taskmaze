package cmd

import (
	"fmt"
	"strings"
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

func validateDeadLine(deadline string) (string, error) {
	deadline = strings.ToLower(deadline)

	if deadline == "today" {
		return time.Now().Format("2006-01-02"), nil
	}
	if deadline == "tomorrow" {
		return time.Now().AddDate(0, 0, 1).Format("2006-01-02"), nil
	}
	for _, day := range []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"} {
		if deadline == day {
			return getNextDayOfWeek(time.Monday).Format("2006-01-02"), nil
		}
	}

	date, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return "" , err
	}
	if date.Before(time.Now()) {
		return "", fmt.Errorf("Deadline should be greater than now")
	}else{
		return date.Format("2006-01-02"), nil
	}
}


func getNextDayOfWeek(day time.Weekday) time.Time {
	today := time.Now().Weekday()
	if today == day {
		return time.Now().AddDate(0, 0, 7)
	}
	if today > day {
		return time.Now().AddDate(0, 0, 7-int(today-day))
	}
	return time.Now().AddDate(0, 0, int(day-today))
}


// valid value for reminder:
// remind_me_by
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
	if !validateDeadLine(dueDate) {
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
