package dao

import "gorm.io/gorm"

type BaseGormDAO struct {
	Session *gorm.DB
}
