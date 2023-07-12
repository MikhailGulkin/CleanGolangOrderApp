package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo/product"
	"regexp"
)

type Product struct {
	product.ProductID
	Price        product.ProductPrice
	Name         string
	Discount     product.ProductDiscount
	Quantity     int32
	Description  string
	Availability bool
}

func (Product) Create(
	productID product.ProductID,
	price product.ProductPrice,
	discount product.ProductDiscount,
	description string,
	name string,
) (Product, error) {
	return Product{
		ProductID:    productID,
		Price:        price,
		Name:         name,
		Discount:     discount,
		Quantity:     1,
		Description:  description,
		Availability: true,
	}, nil
}
func (product *Product) UpdateName(name string) error {
	matched, err := regexp.MatchString("^[A-Z].*", name)
	if err != nil {
		return err
	}
	if !matched {
		exception := exceptions.InvalidProductNameUpdate{}.Exception(name)
		return &exception
	}
	product.Name = name
	return nil
}
