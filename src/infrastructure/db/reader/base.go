package reader

import "gorm.io/gorm"

type BaseGormDAO struct {
	Session *gorm.DB
}
