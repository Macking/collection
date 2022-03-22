package minio

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDBConnection = ""

func DBOperate() *gorm.DB {
	dsn := "root:root@tcp(192.168.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
