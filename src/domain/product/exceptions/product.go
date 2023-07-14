package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common"
)

type InvalidPriceProductCreation struct {
	common.CustomException
}
type InvalidDiscountProductCreation struct {
	common.CustomException
}
type InvalidProductNameUpdate struct {
	common.CustomException
}

func (e InvalidPriceProductCreation) Exception(context string) InvalidPriceProductCreation {
	return InvalidPriceProductCreation{
		CustomException: common.CustomException{
			Message: "Price cannot be negative;",
			Ctx:     fmt.Sprintf("price: `%s`", context),
		}}
}
func (e InvalidDiscountProductCreation) Exception(context string) InvalidDiscountProductCreation {
	return InvalidDiscountProductCreation{
		CustomException: common.CustomException{
			Message: "Discount must be in from 0 to 99;",
			Ctx:     fmt.Sprintf("discount: `%s`", context),
		}}
}
func (e InvalidProductNameUpdate) Exception(context string) InvalidProductNameUpdate {
	return InvalidProductNameUpdate{
		CustomException: common.CustomException{
			Message: "Product name must be start with capital letter",
			Ctx:     fmt.Sprintf("api name: `%s`", context),
		}}
}
