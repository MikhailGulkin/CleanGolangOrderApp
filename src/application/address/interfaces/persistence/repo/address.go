package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/address/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
)

type AddressRepo interface {
	AcquireAddressByID(addressID id.ID) (aggregate.Address, error)
	AddAddress(address aggregate.Address, tx interface{}) error
}
