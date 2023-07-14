package utils

import (
	"bytes"
	"encoding/json"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"github.com/google/uuid"
	"log"
)

func CreateValidOrderCommand(model models.Order) command.CreateOrderCommand {
	return command.CreateOrderCommand{
		OrderID:         model.ID,
		PaymentMethod:   string(consts.Cash),
		ProductsIDs:     []uuid.UUID{model.Products[0].ID},
		UserID:          model.Client.ID,
		DeliveryAddress: model.AddressID,
	}
}

func CreateValidByteOrder(model models.Order) *bytes.Reader {
	marshalled, err := json.Marshal(CreateValidOrderCommand(model))
	if err != nil {
		log.Fatalf("impossible to marshall api: %s", err)
	}
	return bytes.NewReader(marshalled)
}
