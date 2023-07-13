package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain"
)

type OrderIsClosed struct {
	domain.CustomException
}

func (e OrderIsClosed) Exception(context string) OrderIsClosed {
	return OrderIsClosed{
		CustomException: domain.CustomException{
			Message: "Order already closed;",
			Ctx:     fmt.Sprintf("order id: `%s`", context),
		}}
}

type InvalidSerialNumber struct {
	domain.CustomException
}

func (e InvalidSerialNumber) Exception(context string) InvalidSerialNumber {
	return InvalidSerialNumber{
		CustomException: domain.CustomException{
			Message: "Invalid serial number;",
			Ctx:     fmt.Sprintf("serial number `%s`", context),
		}}
}

type ProductAlreadyContained struct {
	domain.CustomException
}

func (e ProductAlreadyContained) Exception(productID string, context string) ProductAlreadyContained {
	return ProductAlreadyContained{
		CustomException: domain.CustomException{
			Message: fmt.Sprintf("Product with this id %s already contained;", productID),
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}

type OrderProductNotExists struct {
	domain.CustomException
}

func (e OrderProductNotExists) Exception(productID string, context string) OrderProductNotExists {
	return OrderProductNotExists{
		CustomException: domain.CustomException{
			Message: fmt.Sprintf("Product with this id %s not exist in order;", productID),
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}

type OrderProductsEmpty struct {
	domain.CustomException
}

func (e OrderProductsEmpty) Exception(context string) OrderProductsEmpty {
	return OrderProductsEmpty{
		CustomException: domain.CustomException{
			Message: "Order cannot doesn't have products;",
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}
