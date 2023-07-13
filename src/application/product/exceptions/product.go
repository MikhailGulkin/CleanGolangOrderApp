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

type ProductIDNotExist struct {
	domain.CustomException
}

func (e ProductIDNotExist) Exception(context string) ProductIDNotExist {
	return ProductIDNotExist{
		CustomException: domain.CustomException{
			Message: "Product with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}

type ProductIDsNotExist struct {
	domain.CustomException
}

func (e ProductIDsNotExist) Exception(context []string) ProductIDsNotExist {
	return ProductIDsNotExist{
		CustomException: domain.CustomException{
			Message: "Product with this ids not exist;",
			Ctx:     fmt.Sprintf("ids: `%v`", context),
		}}
}
