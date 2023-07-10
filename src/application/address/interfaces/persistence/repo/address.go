package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/address"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type AddressRepo interface {
	AcquireAddressByID(addressID vo.AddressID) (address.Address, error)
}
