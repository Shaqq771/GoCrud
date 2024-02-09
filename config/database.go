package config

import (
	//"fmt"
	"gocrud/model"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var DB = "root:root@tcp(127.0.0.1:3306)/GOCRUD?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func ConnectDB() {
	connectionString := "root:root@tcp(127.0.0.1:3307)/GOCRUD?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Alamat{}, &model.User{})


	DB = db

	log.Println("Database Connected")
}