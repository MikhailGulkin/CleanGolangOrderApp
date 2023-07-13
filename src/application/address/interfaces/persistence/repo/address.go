package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
)

type AddressRepo interface {
	AcquireAddressByID(addressID id.ID) (address.Address, error)
	AddAddress(address address.Address, tx interface{}) error
}
