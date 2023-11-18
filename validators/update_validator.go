package validators

import (
	"time"

	constant "github.com/DrAnonymousNet/taskmaze/utils"
)


func ValidateUpdateArgs(updateData map[string]interface{}) (bool, error) {
	if title, exist := updateData[constant.TITLE]; exist {
		if valid, err := validateTitle(title.(string)); !valid {
			return false, err
		}
	}
	deadline, deadlineExist := updateData[constant.DEADLINE]
	remindMe, remindMeExist := updateData[constant.REMIND_ME] 
	update := true 

	if deadlineExist {
		if valid, err := validateDeadLine(deadline.(time.Time), update); !valid {
			return false, err
		}
	}
	if remindMeExist{
		if valid, err := validateReminder(remindMe.(time.Time) , deadline.(time.Time)); !valid{
			return false, err
		}
		}else{
		// Just a check in place
		panic("The remindMe and deadline should be passed in together")
		}


	if priority, priorityExist := updateData[constant.PRIORITY]; priorityExist{
		if valid, err := validatePriority(priority.(string)); valid{
			return false, err
		}
	} 


	return true, nil
}