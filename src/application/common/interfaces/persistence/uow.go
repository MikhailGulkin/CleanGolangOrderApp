package persistence

type UoW interface {
	StartTx()
	GetTx() interface{}
	Commit()
	Rollback()
}
