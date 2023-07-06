package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
)

type ProductNameNotExist struct {
	domain.CustomException
}

func (e ProductNameNotExist) Exception(context string) ProductNameNotExist {
	return ProductNameNotExist{
		CustomException: domain.CustomException{
			Message: "Product with this name not exist;",
			Ctx:     fmt.Sprintf("name `%s`", context),
		}}
}
