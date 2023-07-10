package aggregate

import (
	domain2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/address"
	"github.com/google/uuid"
	"testing"
)

func TestIncorrectSerialOrderCreate(t *testing.T) {
	var incorrectSerial = 110
	_, err := domain2.Order.Create(domain2.Order{}, uuid.New(), &[]domain.Product{}, address.Address{}, incorrectSerial)
	if err == nil {
		t.Error("Wrong Order created with serial: ", incorrectSerial)
	}
}
