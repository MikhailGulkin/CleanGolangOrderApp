package uow

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"gorm.io/gorm"
)

type GormUoW struct {
	persistence.UoW
	Session *gorm.DB
	tx      *gorm.DB
}

func (uow *GormUoW) StartTx() any {
	uow.tx = uow.Session.Begin()
	return uow.tx
}

func (uow *GormUoW) GetTx() any {
	if uow.tx != nil {
		return uow.tx
	}
	uow.tx = uow.Session.Begin()
	return uow.tx
}
func (uow *GormUoW) Commit() error {
	uow.tx.Commit()
	uow.tx = nil
	return nil
}
func (uow *GormUoW) Rollback() {
	uow.tx.Rollback()
	uow.tx = nil
}
