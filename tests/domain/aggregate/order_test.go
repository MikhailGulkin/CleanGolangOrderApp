package aggregate

import (
	domain2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/google/uuid"
	"testing"
)

func TestIncorrectSerialOrderCreate(t *testing.T) {
	var incorrectSerial = 110
	_, err := domain2.Order.Create(domain2.Order{}, vo.OrderID{Value: uuid.New()}, order.OrderAddress{}, order.OrderClient{}, incorrectSerial)
	if err == nil {
		t.Error("Wrong Order created with serial: ", incorrectSerial)
	}
}
