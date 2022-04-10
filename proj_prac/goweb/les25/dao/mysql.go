package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMYSQL() (err error) {

	DB, err = gorm.Open("mysql", "root:111111@(hadoop102:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		return err
	}

	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
