package validators

import (
	"fmt"
	"strings"
	"time"
)

func validateTitle(title string) (bool, error) {
	if strings.Contains(title, "|") {
		return false, fmt.Errorf("title should not contain the pipe character '|'")
	}
	return true, nil
}

func validateDeadLine(deadline time.Time, update bool) (bool, error) {

	if deadline.IsZero() && !update {
		return false, fmt.Errorf("deadline should not be empty")
	}
	if deadline.IsZero() && update{
		return true, nil
	}
	if deadline.Before(time.Now()) && !update {
		return false, fmt.Errorf("deadline should be in the future")
	} else {
		return true, nil
	}
}

func validateReminder(reminder time.Time, deadline time.Time) (bool, error) {
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
	update bool,
) (bool, error) {

	if valid, err := validateDeadLine(deadline, update); !valid {
		return false, err
	}

	if !remindMe.IsZero() {
		if valid, err := validateReminder(remindMe, deadline); !valid {
			return false, err
		}
	}

	if valid, err := validatePriority(priority); !valid {
		return false, err
	}

	if valid, err := validateTitle(title); !valid{
		return false, err
	}
	return true, nil
}
