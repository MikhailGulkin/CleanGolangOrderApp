package aggregate

import (
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/vo"
	"github.com/google/uuid"
	"testing"
)

func TestIncorrectSerialOrderCreate(t *testing.T) {
	var incorrectSerial = 110
	_, err := domain.Order.Create(domain.Order{}, vo.OrderID{Value: uuid.New()}, order.OrderAddress{}, order.OrderClient{}, incorrectSerial)
	if err == nil {
		t.Error("Wrong Order created with serial: ", incorrectSerial)
	}
}
