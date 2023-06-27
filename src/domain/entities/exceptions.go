package domain

import "github.com/MikhailGulkin/simpleGoOrderApp/src/domain"

type InvalidPriceProductCreation struct {
	domain.CustomException
}
type InvalidDiscountProductCreation struct {
	domain.CustomException
}

func (e InvalidPriceProductCreation) Exception(context string) InvalidPriceProductCreation {
	return InvalidPriceProductCreation{
		CustomException: domain.CustomException{
			Message: "Price cannot be negative; price:",
			Ctx:     context,
		}}
}
func (e InvalidDiscountProductCreation) Exception(context string) InvalidDiscountProductCreation {
	return InvalidDiscountProductCreation{
		CustomException: domain.CustomException{
			Message: "Discount must be in from 0 to 99; discount:",
			Ctx:     context,
		}}
}
