package vo

type CustomerBalance struct {
	AvailableMoney BalanceMoney
	FrozenMoney    BalanceMoney
	TotalMoney     float64
}

func (CustomerBalance) Create() CustomerBalance {
	return CustomerBalance{
		AvailableMoney: BalanceMoney{Value: 0},
		FrozenMoney:    BalanceMoney{Value: 0},
		TotalMoney:     0,
	}
}
