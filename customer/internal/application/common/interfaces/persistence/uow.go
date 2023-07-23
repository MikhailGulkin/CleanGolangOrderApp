package persistence

type UoW interface {
	StartTx() any
	GetTx() any
	Commit() error
	Rollback()
}
