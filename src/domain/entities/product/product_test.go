package product

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	"testing"
)

func TestIncorrectPriceProductCreate(t *testing.T) {
	var incorrectPrice float64 = -100

	_, err := Product.Create(Product{}, incorrectPrice, 0, "", "")
	var priceError *domain.InvalidPriceProductCreation
	if !errors.As(err, &priceError) {
		t.Error("Product with negative price created, price: ", incorrectPrice)
	}
}
func TestIncorrectDiscountProductCreate(t *testing.T) {
	var discount int32 = 3000
	_, err := Product.Create(Product{}, 100, discount, "", "")
	var discountError *domain.InvalidDiscountProductCreation
	if !errors.As(err, &discountError) {
		t.Error("Product with unlivable discount created, price: ", discount)
	}
}
