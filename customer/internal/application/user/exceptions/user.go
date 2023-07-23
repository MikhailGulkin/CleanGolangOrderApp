package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common"
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

type UserAddressIDNotExist struct {
	common.CustomException
}

func (e UserAddressIDNotExist) Exception(context string) UserAddressIDNotExist {
	return UserAddressIDNotExist{
		CustomException: common.CustomException{
			Message: "User address with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
