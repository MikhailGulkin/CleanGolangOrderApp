package consts

type TransactionType string

const (
	DEPOSIT            TransactionType = "deposit"
	PURCHASE_PENDING   TransactionType = "purchase_pending"
	PURCHASE_COMPLETED TransactionType = "purchase_completed"
	PURCHASE_FAILED    TransactionType = "purchase_failed"
)
