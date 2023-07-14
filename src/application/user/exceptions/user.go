package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common"
)

type UserIDNotExist struct {
	common.CustomException
}

func (e UserIDNotExist) Exception(context string) UserIDNotExist {
	return UserIDNotExist{
		CustomException: common.CustomException{
			Message: "User with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
