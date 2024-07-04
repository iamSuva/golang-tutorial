package config

import (
	"fmt"
	"github.com/iamSuva/postgres_restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("failed to connect to database")
	}
    db.AutoMigrate(&models.User{})
	DB=db
	fmt.Println("connected to databse")
}