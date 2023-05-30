package dao

import (
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func Init() {
	var err error
	log.Println(os.Getenv("MYSQL_USERNAME"))
	dsn := os.Getenv("MYSQL_USERNAME") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_IP") + ":" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DATABASE")
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	log.Println("MySQL has connected!")
}
