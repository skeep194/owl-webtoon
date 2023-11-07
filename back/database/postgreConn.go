package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgreDB *gorm.DB = nil

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST_NAME"), os.Getenv("DB_USER_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	var err error
	PostgreDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("db connection fail")
	}
}
