package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DbCon *gorm.DB
)

func ConnectToDatabase() {
	connection, err := gorm.Open(mysql.Open(returnStringConnection()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DbCon = connection
}

func returnStringConnection() string {

	charset := os.Getenv("DB_CHARSET")
	if charset == "" {
		charset = "utf8mb4"
	}

	parseTime := os.Getenv("DB_PARSETIME")
	if parseTime == "" {
		parseTime = "True"
	}

	loc := os.Getenv("DB_LOC")
	if loc == "" {
		loc = "Local"
	}

	config := []interface{}{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		charset,
		parseTime,
		loc,
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", config...)
}
