package vo

type Balance struct {
	AvailableMoney Money
	FrozenMoney    Money
}

func NewCustomerBalance() Balance {
	return Balance{
		AvailableMoney: Money{Value: 0},
		FrozenMoney:    Money{Value: 0},
	}
}

func (b *Balance) SubAvailableMoney(money Money) {
	b.AvailableMoney.Value -= money.Value
}
func (b *Balance) AddFrozenMoney(money Money) {
	b.FrozenMoney.Value += money.Value
}
func (b *Balance) EqAvailableMoney(money Money) bool {
	if b.AvailableMoney.Value == money.Value {
		return false
	}
	return true
}
func (b *Balance) AddAvailableMoney(money Money) {
	b.AvailableMoney.Value += money.Value
}
func (b *Balance) SubFrozenMoney(money Money) {
	b.FrozenMoney.Value -= money.Value
}
func (b *Balance) AddFrozenMoneyFromAvailable(money Money) {
	b.AvailableMoney.Value -= money.Value
	b.FrozenMoney.Value += money.Value
}

func (b *Balance) DepositBalance(money Money) Balance {
	b.AddAvailableMoney(money)
	return *b
}
func (b *Balance) Purchase(money Money) Balance {
	b.AddFrozenMoneyFromAvailable(money)
	return *b
}
