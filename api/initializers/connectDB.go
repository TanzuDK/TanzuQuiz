package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	instancename := os.Getenv("instancename")
	dbname := os.Getenv("dbname")
	username := os.Getenv("username")
	password := os.Getenv("password")
	port := "5432"

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Copenhagen", instancename, username, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("For trubleshooting - Printing connection details")
		log.Println("instancename: " + instancename)
		log.Println("dbname: " + dbname)
		log.Println("username: " + username)
		log.Println("password: " + password)
		log.Println("port: " + port)
		log.Fatal("Failed to connect to the Database")

	}
	fmt.Println("? Connected Successfully to the Database")
}
