package persistence

type UoW interface {
	StartTx()
	GetTx() any
	Commit() error
	Rollback()
}
