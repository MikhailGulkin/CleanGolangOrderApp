package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common"
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

type ProductIDsNotExist struct {
	common.CustomException
}

func (e ProductIDsNotExist) Exception(context []string) ProductIDsNotExist {
	return ProductIDsNotExist{
		CustomException: common.CustomException{
			Message: "Product with this ids not exist;",
			Ctx:     fmt.Sprintf("ids: `%v`", context),
		}}
}
