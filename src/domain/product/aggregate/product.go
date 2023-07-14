package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/exceptions"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/vo"
	"regexp"
)

type Product struct {
	vo.ProductID
	Price        vo.ProductPrice
	Name         string
	Discount     vo.ProductDiscount
	Quantity     int32
	Description  string
	Availability bool
}

func (Product) Create(
	productID vo.ProductID,
	price vo.ProductPrice,
	discount vo.ProductDiscount,
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
