package persistence

type UoW interface {
	StartTx()
	GetTx() interface{}
	Commit() error
	Rollback() error
}
