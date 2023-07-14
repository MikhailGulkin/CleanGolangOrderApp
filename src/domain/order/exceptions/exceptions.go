package exceptions

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common"
)

type OrderIsClosed struct {
	common.CustomException
}

func (e OrderIsClosed) Exception(context string) OrderIsClosed {
	return OrderIsClosed{
		CustomException: common.CustomException{
			Message: "Order already closed;",
			Ctx:     fmt.Sprintf("order id: `%s`", context),
		}}
}

type InvalidSerialNumber struct {
	common.CustomException
}

func (e InvalidSerialNumber) Exception(context string) InvalidSerialNumber {
	return InvalidSerialNumber{
		CustomException: common.CustomException{
			Message: "Invalid serial number;",
			Ctx:     fmt.Sprintf("serial number `%s`", context),
		}}
}

type ProductAlreadyContained struct {
	common.CustomException
}

func (e ProductAlreadyContained) Exception(productID string, context string) ProductAlreadyContained {
	return ProductAlreadyContained{
		CustomException: common.CustomException{
			Message: fmt.Sprintf("Product with this id %s already contained;", productID),
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}

type OrderProductNotExists struct {
	common.CustomException
}

func (e OrderProductNotExists) Exception(productID string, context string) OrderProductNotExists {
	return OrderProductNotExists{
		CustomException: common.CustomException{
			Message: fmt.Sprintf("Product with this id %s not exist in order;", productID),
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}

type OrderProductsEmpty struct {
	common.CustomException
}

func (e OrderProductsEmpty) Exception(context string) OrderProductsEmpty {
	return OrderProductsEmpty{
		CustomException: common.CustomException{
			Message: "Order cannot doesn't have products;",
			Ctx:     fmt.Sprintf("order id `%s`", context),
		}}
}
