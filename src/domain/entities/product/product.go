package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
	domain2 "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	vo "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
	"github.com/google/uuid"
	"strconv"
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
		idError := domain.InvalidUUIDCreation.Exception(domain.InvalidUUIDCreation{}, "Product Creation failure", err.Error())
		return Product{}, &idError
	}
	if price < 0 {
		priceError := domain2.InvalidPriceProductCreation.Exception(domain2.InvalidPriceProductCreation{}, strconv.Itoa(int(price)))
		return Product{}, &priceError
	}

	if discount < 0 || discount > 99 {
		discountError := domain2.InvalidDiscountProductCreation.Exception(domain2.InvalidDiscountProductCreation{}, strconv.Itoa(int(discount)))
		return Product{}, &discountError
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
