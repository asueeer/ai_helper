package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", "postgres://nearby:nearby123@47.104.186.111/nearby_boe_test?sslmode=disable")
	if err != nil {
		log.Println("数据库连接失败" + err.Error())
		//panic("数据库连接失败" + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
