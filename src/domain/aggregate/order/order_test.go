package domain

import (
	domain "github.com/MikhailGulkin/simpleOrderApp/src/domain/entities"
	"testing"
)

func TestIncorrectSerialOrderCreate(t *testing.T) {
	var incorrectSerial int = 110
	_, err := Order.create(Order{}, &[]domain.Product{}, incorrectSerial)
	if err == nil {
		t.Error("Wrong Order created with serial: ", incorrectSerial)
	}
}
