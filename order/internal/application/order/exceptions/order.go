package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common"
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

type OrderClientIDNotExist struct {
	common.CustomException
}

func (e OrderClientIDNotExist) Exception(context string) OrderClientIDNotExist {
	return OrderClientIDNotExist{
		CustomException: common.CustomException{
			Message: "Client with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}

type OrderAddressIDNotExist struct {
	common.CustomException
}

func (e OrderAddressIDNotExist) Exception(context string) OrderAddressIDNotExist {
	return OrderAddressIDNotExist{
		CustomException: common.CustomException{
			Message: "Address with this id not exist;",
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
