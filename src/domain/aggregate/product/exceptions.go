package product

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
)

type InvalidPriceProductCreation struct {
	domain.CustomException
}
type InvalidDiscountProductCreation struct {
	domain.CustomException
}
type InvalidProductNameUpdate struct {
	domain.CustomException
}

func (e InvalidPriceProductCreation) Exception(context string) InvalidPriceProductCreation {
	return InvalidPriceProductCreation{
		CustomException: domain.CustomException{
			Message: "Price cannot be negative;",
			Ctx:     fmt.Sprintf("price: `%s`", context),
		}}
}
func (e InvalidDiscountProductCreation) Exception(context string) InvalidDiscountProductCreation {
	return InvalidDiscountProductCreation{
		CustomException: domain.CustomException{
			Message: "Discount must be in from 0 to 99;",
			Ctx:     fmt.Sprintf("discount: `%s`", context),
		}}
}
func (e InvalidProductNameUpdate) Exception(context string) InvalidProductNameUpdate {
	return InvalidProductNameUpdate{
		CustomException: domain.CustomException{
			Message: "Product name must be start with capital letter",
			Ctx:     fmt.Sprintf("product name: `%s`", context),
		}}
}
