package database

import (
	"fmt"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigDB() {
	var err error
	const finalMYSQL = "root:@tcp(127.0.0.1:3306)/examplego?charset=utf8mb4&parseTime=True&loc=Local"
	// const finalPostgres = "postgresql://postgres:xxx@localhost:5432/examplego?sslmode=disable&TimeZone=Asia/Jakarta"

	dsn := finalMYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to finalMYSQL")
	}
	fmt.Println("Connected to database finalMYSQL")

}
