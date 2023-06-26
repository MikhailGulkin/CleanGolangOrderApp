package domain

import (
	"errors"
	vo "github.com/MikhailGulkin/simpleOrderApp/src/domain/value_object"
	"github.com/google/uuid"
)

type Product struct {
	vo.ProductId
	price        float64
	discount     int32
	quantity     int32
	description  string
	availability bool
}

func (_ Product) create(price float64, discount int32, description string) (Product, error) {
	productId, err := uuid.NewUUID()

	if err != nil {
		return Product{}, errors.New("When ProductId was created error occurred: " + err.Error())
	}
	if price < 0 {
		return Product{}, errors.New("price cannot be negative")
	}

	if discount < 0 || discount > 99 {
		return Product{}, errors.New("discount must be in from 0 to 99")
	}
	return Product{
		ProductId:    vo.ProductId{Value: productId},
		price:        price,
		discount:     discount,
		quantity:     1,
		description:  description,
		availability: true,
	}, nil
}
