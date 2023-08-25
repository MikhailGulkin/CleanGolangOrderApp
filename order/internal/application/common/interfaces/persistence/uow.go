package persistence

type UoW interface {
	Get() UoW
	StartTx() any
	GetTx() any
	Commit() error
	Rollback()
}
