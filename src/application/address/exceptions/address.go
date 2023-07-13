package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
)

type AddressIDNotExist struct {
	domain.CustomException
}

func (e AddressIDNotExist) Exception(context string) AddressIDNotExist {
	return AddressIDNotExist{
		CustomException: domain.CustomException{
			Message: "Address with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
