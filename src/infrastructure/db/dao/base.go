package dao

import "gorm.io/gorm"

type BaseGormDAO struct {
	session *gorm.DB
}
