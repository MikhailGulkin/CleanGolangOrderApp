package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/google/uuid"
	"regexp"
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

func (Product) Create(productID uuid.UUID, price float64, discount int32, description string, name string) (Product, error) {
	if price < 0 {
		priceError := InvalidPriceProductCreation{}.Exception(strconv.Itoa(int(price)))
		return Product{}, &priceError
	}

	if discount < 0 || discount > 99 {
		discountError := InvalidDiscountProductCreation{}.Exception(strconv.Itoa(int(discount)))
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
func (product *Product) UpdateName(name string) error {
	matched, err := regexp.MatchString("^[A-Z].*", name)
	if err != nil {
		return err
	}
	if !matched {
		exception := InvalidProductNameUpdate{}.Exception(name)
		return &exception
	}
	product.Name = name
	return nil
}
