package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common"
)

type AddressIDNotExist struct {
	common.CustomException
}

func (e AddressIDNotExist) Exception(context string) AddressIDNotExist {
	return AddressIDNotExist{
		CustomException: common.CustomException{
			Message: "Address with this id not exist;",
			Ctx:     fmt.Sprintf("id `%s`", context),
		}}
}
