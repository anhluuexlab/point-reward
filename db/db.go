package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	uri := "root:password@tcp(127.0.0.1:5123)/_f?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
