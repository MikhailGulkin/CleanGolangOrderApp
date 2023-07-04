package aggregate

import (
	domain2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/address"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"testing"
)

func TestIncorrectSerialOrderCreate(t *testing.T) {
	var incorrectSerial = 110
	_, err := domain2.Order.Create(domain2.Order{}, &[]domain.Product{}, address.Address{}, incorrectSerial)
	if err == nil {
		t.Error("Wrong Order created with serial: ", incorrectSerial)
	}
}
