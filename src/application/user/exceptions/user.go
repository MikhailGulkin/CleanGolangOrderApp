package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
)

type UserIDNotExist struct {
	domain.CustomException
}

func (e UserIDNotExist) Exception(context string) UserIDNotExist {
	return UserIDNotExist{
		CustomException: domain.CustomException{
			Message: "User with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
