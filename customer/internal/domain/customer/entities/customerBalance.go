package entities

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/customer/vo"
	"github.com/google/uuid"
)

type CustomerBalance struct {
	BalanceID      uuid.UUID
	AvailableMoney vo.BalanceMoney
	FrozenMoney    vo.BalanceMoney
	TotalMoney     float64
}

func (CustomerBalance) Create() CustomerBalance {
	return CustomerBalance{
		BalanceID:      uuid.New(),
		AvailableMoney: vo.BalanceMoney{Value: 0},
		FrozenMoney:    vo.BalanceMoney{Value: 0},
		TotalMoney:     0,
	}
}
