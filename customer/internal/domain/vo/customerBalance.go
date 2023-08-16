package vo

type CustomerBalance struct {
	AvailableMoney Money
	FrozenMoney    Money
}

func NewCustomerBalance() CustomerBalance {
	return CustomerBalance{
		AvailableMoney: Money{Value: 0},
		FrozenMoney:    Money{Value: 0},
	}
}

func (c *CustomerBalance) SubAvailableMoney(money Money) {
	c.AvailableMoney.Value -= money.Value
}
func (c *CustomerBalance) AddFrozenMoney(money Money) {
	c.FrozenMoney.Value += money.Value
}
func (c *CustomerBalance) EqAvailableMoney(money Money) bool {
	if c.AvailableMoney.Value == money.Value {
		return false
	}
	return true
}
