package product

import (
	"bytes"
	"encoding/json"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"log"
)

func CreateValidProductCommand() command.CreateProductCommand {
	return command.CreateProductCommand{
		Price:       100,
		Discount:    10,
		Description: "Some Product",
		Name:        "Some name",
	}
}

func CreateValidByteProduct() *bytes.Reader {
	marshalled, err := json.Marshal(CreateValidProductCommand())
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}
	return bytes.NewReader(marshalled)
}
func CreateInvalidDiscountByteProduct() *bytes.Reader {
	product := CreateValidProductCommand()
	product.Discount = -100
	marshalled, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}
	return bytes.NewReader(marshalled)
}
func CreateInvalidPriceByteProduct() *bytes.Reader {
	product := CreateValidProductCommand()
	product.Price = -100
	marshalled, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}
	return bytes.NewReader(marshalled)
}