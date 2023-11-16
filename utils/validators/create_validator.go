package validators

import (
	"fmt"
	"time"
)



func validateDeadLine(deadline time.Time) (bool, error) {
	if deadline.IsZero() {
		return false, fmt.Errorf("deadline should not be empty")
	}
	if deadline.Before(time.Now()) {
		return false, fmt.Errorf("deadline should be in the future")
	}else{
		return true, nil
	}
}




func validateReminder(reminder time.Time, deadline time.Time) (bool, error) {
	// The reminder + the current time should not be greater that the dueTime
	if reminder.After(deadline) {
		return false, fmt.Errorf("reminder should be before the deadline")
	}
	return true, nil

}

func validatePriority(priority string) (bool, error) {
	if priority == "high" || priority == "medium" || priority == "low" {
		return true, nil
	}
	return false, fmt.Errorf("priority should be high, medium or low")
}


func ValidateCreateArgs(
	title string,
	deadline time.Time,
	priority string,
	remindMe time.Time,
) (bool, error) {

	if valid, err := validateDeadLine(deadline); !valid  {
		return false, err
	}

	if !remindMe.IsZero(){
		if valid, err := validateReminder(remindMe, deadline); !valid {
			return false, err
		}
	}

	if valid, err := validatePriority(priority); !valid {
		return false, err
	}
	return true, nil
}
