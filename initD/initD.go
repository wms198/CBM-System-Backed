package initD

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//var dsn = "root:@tcp(localhost)/cbmSystem?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect mysql DB")
	}
}
