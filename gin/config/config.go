package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open("debian-sys-maint:nFVTfJnpqDmlmCSO@tcp(localhost:3306)/aobi?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("无法连接至数据库")
	}
}
