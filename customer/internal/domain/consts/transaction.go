package consts

type TransactionType string

const (
	DEPOSIT  TransactionType = "deposit"
	PURCHASE TransactionType = "purchase"
	FREZEE   TransactionType = "freeze"
	UNFREZEE TransactionType = "unfreeze"
)
