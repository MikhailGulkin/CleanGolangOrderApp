package consts

type TransactionType string

const (
	WRITE_OFF TransactionType = "write_off"
	REFILLING TransactionType = "refilling"
	FREZEE    TransactionType = "freeze"
	UNFREZEE  TransactionType = "unfreeze"
)
