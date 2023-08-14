package vo

type BalanceMoney struct {
	Value float64
}

func (BalanceMoney) Create(money float64) (BalanceMoney, error) {
	if money < 0 {
		// TODO: add error handling
		return BalanceMoney{}, nil
	}
	return BalanceMoney{Value: money}, nil
}
