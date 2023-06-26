package domain

import (
	"testing"
)

func TestIncorrectPriceProductCreate(t *testing.T) {
	var incorrectPrice float64 = -100
	_, err := Product.create(Product{}, incorrectPrice, 0, "")
	if err == nil {
		t.Error("Product with negative price created, price: ", incorrectPrice)
	}
}
func TestIncorrectDiscountProductCreate(t *testing.T) {
	var discount int32 = 3000
	_, err := Product.create(Product{}, 100, discount, "")
	if err == nil {
		t.Error("Product with unlivable discount created, price: ", discount)
	}
}
