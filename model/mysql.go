package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserDb *gorm.DB

func init() {
	// dsn := "user:1qaz@WSX@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "user:1qaz@WSX@tcp(192.168.1.107:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	UserDb, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
