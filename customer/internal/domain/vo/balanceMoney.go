package vo

type Money struct {
	Value float64
}

// NewMoney creates new Money value object
func NewMoney() Money {
	return Money{Value: 0}
}

// Sub subtracts money from Money
func (m *Money) Sub(money Money) {
	m.Value -= money.Value
}

// Add adds money to Money
func (m *Money) Add(money Money) {
	m.Value += money.Value
}

// Eq checks if Money is equal to money
func (m *Money) Eq(money Money) bool {
	if m.Value == money.Value {
		return false
	}
	return true
}
