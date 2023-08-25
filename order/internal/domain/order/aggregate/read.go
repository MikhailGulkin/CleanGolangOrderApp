package aggregate

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/exceptions"
	"github.com/google/uuid"
	"strconv"
)

func getCurrentSerialNumber(serialNumber int) (int, error) {
	if serialNumber > 100 || serialNumber < 0 {
		exception := exceptions.InvalidSerialNumber{}.Exception(strconv.Itoa(serialNumber))
		return -1, &exception
	}
	if serialNumber == 100 {
		return 1, nil
	}
	return serialNumber + 1, nil
}
func (o *Order) GetTotalPrice() PriceOrder {
	var totalPrice float64
	var totalDiscount float64
	for _, orderProduct := range o.Products {
		totalPrice += orderProduct.Price
		totalDiscount += float64(orderProduct.Discount)
	}
	o.totalPrice = PriceOrder(totalPrice - ((totalDiscount / 100) * 100))
	return o.totalPrice
}
func (o *Order) GetAllProductsIds() []uuid.UUID {
	ids := make([]uuid.UUID, len(o.Products))
	for index, product := range o.Products {
		ids[index] = product.ProductID
	}
	return ids
}
