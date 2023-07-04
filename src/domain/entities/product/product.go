package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
	domainEntity "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities"
	vo "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
	"github.com/google/uuid"
	"strconv"
)

type Product struct {
	vo.ProductID
	Price        float64
	Name         string
	Discount     int32
	Quantity     int32
	Description  string
	Availability bool
}

func (Product) Create(price float64, discount int32, description string, name string) (Product, error) {
	productID, err := uuid.NewUUID()

	if err != nil {
		idError := domain.InvalidUUIDCreation{}.Exception("Product ID Creation failure", err.Error())
		return Product{}, &idError
	}
	if price < 0 {
		priceError := domainEntity.InvalidPriceProductCreation{}.Exception(strconv.Itoa(int(price)))
		return Product{}, &priceError
	}

	if discount < 0 || discount > 99 {
		discountError := domainEntity.InvalidDiscountProductCreation{}.Exception(strconv.Itoa(int(discount)))
		return Product{}, &discountError
	}
	return Product{
		ProductID:    vo.ProductID{Value: productID},
		Price:        price,
		Name:         name,
		Discount:     discount,
		Quantity:     1,
		Description:  description,
		Availability: true,
	}, nil
}
