package domain

import (
	"errors"
	"github.com/MikhailGulkin/simpleOrderApp/src/domain/entities"
	vo "github.com/MikhailGulkin/simpleOrderApp/src/domain/value_object"
	"github.com/google/uuid"
	"strconv"
)

type Order struct {
	vo.OrderId
	products     []domain.Product
	serialNumber int
}

func (_ Order) create(products *[]domain.Product, previousSerialNumber int) (Order, error) {
	OrderId, idError := uuid.NewUUID()

	if idError != nil {
		return Order{}, errors.New("When OrderId was created error occurred: " + idError.Error())
	}
	serialNumber, serialError := getCurrentSerialNumber(previousSerialNumber)
	if serialError != nil {
		return Order{}, errors.New(serialError.Error())
	}

	return Order{
		OrderId:      vo.OrderId{Value: OrderId},
		products:     *products,
		serialNumber: serialNumber,
	}, nil
}
func getCurrentSerialNumber(serialNumber int) (int, error) {
	if serialNumber > 100 || serialNumber < 1 {
		return -1, errors.New("wrong SerialNumber creation, serial: " + strconv.Itoa(serialNumber))

	}
	if serialNumber == 100 {
		return 1, nil
	}
	return serialNumber + 1, nil
}
