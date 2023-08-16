package vo

type Money struct {
	Value float64
}

func (Money) Create(money float64) (Money, error) {
	if money < 0 {
		// TODO: add error handling
		return Money{}, nil
	}
	return Money{Value: money}, nil
}
