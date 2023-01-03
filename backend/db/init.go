package db

import (
	"fmt"
	"link-shortner/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() {	
	var err error
	// NEEDS TO BE CHANGED TO POSTGRES RUNNING ON DOCKER COMPOSE NETWORK
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Link{})
}