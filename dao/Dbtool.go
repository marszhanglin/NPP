// Dbtool.go
package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db    *gorm.DB
	DEBUG bool
)

func InitDB() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:229575793007@tcp(localhost:3306)/mms?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(10000)
	db.DB().SetConnMaxLifetime(1000)
	db.LogMode(DEBUG)
	return db
}
