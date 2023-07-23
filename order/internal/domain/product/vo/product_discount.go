package vo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/product/exceptions"
	"strconv"
)

type ProductDiscount struct {
	Value int32
}

func (ProductDiscount) Create(discount int32) (ProductDiscount, error) {
	if discount < 0 || discount > 99 {
		discountError := exceptions.InvalidDiscountProductCreation{}.Exception(strconv.Itoa(int(discount)))
		return ProductDiscount{}, &discountError
	}
	return ProductDiscount{Value: discount}, nil
}
