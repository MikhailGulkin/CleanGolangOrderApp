package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/vo"
	"testing"
)

func TestIncorrectPriceProductCreate(t *testing.T) {
	var incorrectPrice float64 = -100
	_, err := vo.ProductPrice{}.Create(incorrectPrice)
	if err != nil {
		return
	}
	var priceError *exceptions.InvalidPriceProductCreation
	if !errors.As(err, &priceError) {
		t.Error("Product with negative price created, price: ", incorrectPrice)
	}
}
func TestIncorrectDiscountProductCreate(t *testing.T) {
	var discount int32 = 3000
	_, err := vo.ProductDiscount{}.Create(discount)
	var discountError *exceptions.InvalidDiscountProductCreation
	if !errors.As(err, &discountError) {
		t.Error("Product with unlivable discount created, price: ", discount)
	}
}
