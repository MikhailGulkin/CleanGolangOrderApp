package product

import (
	"bytes"
	"encoding/json"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"log"
)

func CreateValidProduct() *bytes.Reader {
	product := command.CreateProductCommand{
		Price:       100,
		Discount:    10,
		Description: "Some Product",
		Name:        "Some name",
	}
	marshalled, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}
	return bytes.NewReader(marshalled)
}
