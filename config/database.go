package config

import (
	"fmt"
	"gocrud/model"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := fmt.Sprintf("%v:%v@/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_DATABASE, "%2F")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Alamat{}, &model.User{})


	DB = db

	log.Println("Database Connected")
}