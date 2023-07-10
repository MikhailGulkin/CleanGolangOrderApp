package utils

import (
	"bytes"
	"encoding/json"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/google/uuid"
	"log"
)

func CreateValidProductCommand() command.CreateProductCommand {
	return command.CreateProductCommand{
		ProductID:   uuid.New(),
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
