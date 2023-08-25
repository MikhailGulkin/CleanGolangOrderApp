package uow

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	"gorm.io/gorm"
)

type GormUoW struct {
	persistence.UoW
	Session *gorm.DB
	tx      *gorm.DB
}

func (uow *GormUoW) Get() persistence.UoW {
	return BuildGormUoW(uow.Session)
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
func BuildGormUoW(conn *gorm.DB) persistence.UoW {
	return &GormUoW{Session: conn}
}
