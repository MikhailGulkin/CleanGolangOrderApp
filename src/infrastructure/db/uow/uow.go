package uow

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"gorm.io/gorm"
)

type GormUoW struct {
	persistence.UoW
	Session *gorm.DB
	Tx      *gorm.DB
}

func (uow *GormUoW) StartTx() {
	uow.Tx = uow.Session.Begin()
}

func (uow *GormUoW) GetTx() interface{} {
	if uow.Tx != nil {
		return uow.Tx
	}
	uow.Tx = uow.Session.Begin()
	return uow.Tx
}
func (uow *GormUoW) Commit() {
	uow.Tx.Commit()
}
func (uow *GormUoW) Rollback() {
	uow.Tx.Rollback()
}
