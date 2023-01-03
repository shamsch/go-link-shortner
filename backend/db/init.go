package db

import (
	"fmt"	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

func InitDB() {	
	var err error
	// NEEDS TO BE CHANGED TO POSTGRES RUNNING ON DOCKER COMPOSE NETWORK
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(
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
}