package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	const con = "root:hasansinaga1717@tcp(localhost)/service_manager?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := con
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database not Connected")
	}

	DB = db

	fmt.Println("Database Connected")
}