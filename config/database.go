package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	
	"github.com/lokesh2201013/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=9910994194lokesh dbname=recruitment port=5432 sslmode=disable"
   
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
	if err != nil {
		panic("DB connection failure: " + err.Error())
    }
	log.Println("Connected to database")
	database.AutoMigrate(&models.Profile{})
	database.AutoMigrate(&models.Job{})
	database.AutoMigrate(&models.User{})
    DB = database
}
