package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

type DBDSN string

func NewData(dsn DBDSN, config *gorm.Config) *Data {
	db, err := gorm.Open(mysql.Open(string(dsn)), config)
	if err != nil {
		panic(err)
	}
	return &Data{
		db: db,
	}
}
