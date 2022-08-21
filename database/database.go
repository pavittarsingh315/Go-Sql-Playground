package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{}) // will create an api.db file if one doesn't exist
	if err != nil {
		log.Fatal("Failed to connect to the database!\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connect4ed to the database successfully.")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")

	// TODO: Add migrations

	Database = DbInstance{
		Db: db,
	}
}
