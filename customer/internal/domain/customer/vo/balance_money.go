package vo

import "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/exceptions"

type BalanceMoney struct {
	Value float64
}

func (BalanceMoney) Create(money float64) (BalanceMoney, error) {
	if money < 0 {
		err := exceptions.IncorrectBalance{}.Exception(money)
		return BalanceMoney{}, &err
	}
	return BalanceMoney{Value: money}, nil
}
