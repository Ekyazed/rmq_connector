package database

import (
	"app_connector/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(dbName string) {
	var err error
	dsn := getDsn(dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.FailOnError(err, "Failed to connect database")

	fmt.Printf("Successfully connected to the %s database", dbName)
}

func getDsn(dbName string) string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
