package entities

import (
	"errors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/google/uuid"
	"testing"
)

func TestIncorrectPriceProductCreate(t *testing.T) {
	var incorrectPrice float64 = -100

	_, err := product.Product.Create(product.Product{}, uuid.New(), incorrectPrice, 0, "", "")
	var priceError *product.InvalidPriceProductCreation
	if !errors.As(err, &priceError) {
		t.Error("Product with negative price created, price: ", incorrectPrice)
	}
}
func TestIncorrectDiscountProductCreate(t *testing.T) {
	var discount int32 = 3000
	_, err := product.Product.Create(product.Product{}, uuid.New(), 100, discount, "", "")
	var discountError *product.InvalidDiscountProductCreation
	if !errors.As(err, &discountError) {
		t.Error("Product with unlivable discount created, price: ", discount)
	}
}
