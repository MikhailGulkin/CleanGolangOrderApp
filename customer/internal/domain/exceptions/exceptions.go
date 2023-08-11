package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

type EmptyName struct {
	common.CustomException
}

func (e EmptyName) Exception(message string, context string) EmptyName {
	return EmptyName{
		CustomException: common.CustomException{
			Message: message,
			Ctx:     fmt.Sprintf("name: `%s`", context),
		}}
}

type IncorrectBalance struct {
	common.CustomException
}

func (e IncorrectBalance) Exception(context float64) IncorrectBalance {
	return IncorrectBalance{
		CustomException: common.CustomException{
			Message: "Incorrect balance;",
			Ctx:     fmt.Sprintf("balance: `%f`", context),
		}}
}
