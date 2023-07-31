package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common"
)

type ProductNameNotExist struct {
	common.CustomException
}

func (e ProductNameNotExist) Exception(context string) ProductNameNotExist {
	return ProductNameNotExist{
		CustomException: common.CustomException{
			Message: "Product with this name not exist;",
			Ctx:     fmt.Sprintf("name `%s`", context),
		}}
}

type ProductIDNotExist struct {
	common.CustomException
}

func (e ProductIDNotExist) Exception(context string) ProductIDNotExist {
	return ProductIDNotExist{
		CustomException: common.CustomException{
			Message: "Product with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
