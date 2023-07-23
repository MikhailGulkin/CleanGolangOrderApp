package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/address/aggregate"
)

type AddressRepo interface {
	AddAddress(address aggregate.Address, tx interface{}) error
}
