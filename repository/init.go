package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	// Database Name
	dsn := "root:Excalibur_1551@tcp(127.0.0.1:3306)/998244353?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
