package loaders

import (
	"log"
	"os"

	"github.com/Vikas208/social-media-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil
var err error

func createTables() {
	if DB == nil {
		return
	}
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("Tables created successfully")
}

func ConnectToDatabase() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_NAME")
	log.Println("Connecting to Database")

	dsn := username + ":" + password + "@(" + hostname + ")/" + dbname
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("Database connected successfully")
	createTables()
}
