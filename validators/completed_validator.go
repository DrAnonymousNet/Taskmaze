package validators


func ValidateCompleteArgs(id string)(bool, error){
	return ValidateRetrieveArgs(id)
}