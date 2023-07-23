package address

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/address/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/address/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/address/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"go.uber.org/fx"
)

func NewCreateAddress(repo repo.AddressRepo, uow persistence.UoW) command.CreateAddress {
	return &c.CreateAddressImpl{
		AddressRepo: repo,
		UoW:         uow,
	}
}

var Module = fx.Provide(
	NewCreateAddress,
)
