package validators



func ValidateDeleteArgs(id string) (bool, error) {
	return ValidateRetrieveArgs(id)
}