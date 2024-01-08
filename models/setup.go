// membuat koneksi ke database
package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/belajar_golang_restfull_api"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
