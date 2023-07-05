package conftest

import (
	"fmt"
	"gorm.io/gorm"
)

func CleanTables(conn *gorm.DB) {
	tables := []string{"products"}
	for _, table := range tables {
		conn.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", table))
	}
}
