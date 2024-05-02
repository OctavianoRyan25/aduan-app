package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Connect to database sql
	dsn := "root:root@tcp(localhost:3306)/aduan_app?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migration
}
