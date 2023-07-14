package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common"
)

type OrderIDNotExist struct {
	common.CustomException
}

func (e OrderIDNotExist) Exception(context string) OrderIDNotExist {
	return OrderIDNotExist{
		CustomException: common.CustomException{
			Message: "Order with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
