package validators

import (
	"fmt"
	"strconv"
)


func validateID(id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("id should not be empty")
	}
	// Validate that the id is a valid integer
	if _, err := strconv.Atoi(id); err != nil {
		return false, fmt.Errorf("id should be a valid integer")
	}
	return true, nil
}

func ValidateRetrieveArgs(id string) (bool, error) {
	if valid, err := validateID(id); !valid {
		return false, err
	}
	return true, nil
}