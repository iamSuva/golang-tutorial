package config

import (
	"fmt"
	"github.com/iamSuva/todoapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	dsn := "host=localhost user=suvadip password=suvadip dbname=tododb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("error to connect to database")
		return
	}
	db.AutoMigrate(&models.Todo{})
    DB=db
    fmt.Println("connected to database")
   
}
