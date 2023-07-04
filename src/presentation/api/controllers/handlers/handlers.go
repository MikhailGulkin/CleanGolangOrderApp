package handlers

import (
	interfaces "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProductHandler),
)

func NewProductHandler(interactor interfaces.CreateProduct) ProductHandler {
	return ProductHandler{interactor}
}
