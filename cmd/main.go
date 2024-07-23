package main

import (
	"github.com/ranryl/go-ddd/api"
	"github.com/ranryl/go-ddd/internal/data"
	"gorm.io/gorm"
)

func main() {
	var dsn data.DBDSN = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	app := api.Initialize(dsn, &gorm.Config{})
	app.Run(":8080")
}
