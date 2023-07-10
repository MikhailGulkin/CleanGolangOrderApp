package conftest

import (
	"fmt"
	"gorm.io/gorm"
)

func CleanTables(conn *gorm.DB) {
	var tables []string
	if err := conn.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	fmt.Println(tables)
	for _, table := range tables {
		conn.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", table))
	}
}
