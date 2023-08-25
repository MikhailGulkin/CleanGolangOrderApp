package persistence

type UoW interface {
	Commit() error
	Rollback() error
	Begin() (interface{}, error)
}
type UoWManager interface {
	GetUoW() UoW
}
