package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)



func Connect() {
	godotenv.Load()
	var db_uri string = os.Getenv("DB_URI")
	d, err := gorm.Open("mysql", db_uri)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
