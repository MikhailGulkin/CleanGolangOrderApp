package aggregate

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/entities"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/vo"
	"github.com/google/uuid"
)

type Customer struct {
	vo.UserID
	FullName     vo.FullName
	AddressID    uuid.UUID
	Transactions []entities.CustomerTransactions
	Balance      entities.CustomerBalance
	Orders       []entities.UserOrder
}

func Create(fullName vo.FullName, addressID uuid.UUID) Customer {
	return Customer{
		UserID:    vo.UserID{Value: uuid.New()},
		FullName:  fullName,
		AddressID: addressID,
		Balance:   entities.CustomerBalance{}.Create(),
	}
}

func (c *Customer) AddOrder(order entities.UserOrder) {
	c.Orders = append(c.Orders, order)
}
