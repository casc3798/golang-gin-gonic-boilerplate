package db

import (
	"fmt"
	"os"

	"github.com/CamielaTeam/go-claridad-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Open ...
// Open the database connection and assign the instance to the package scoped var DB
func Open() error {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return err
}

// Automigrate ...
// Applies gorm automigrations
func Automigrate() error {
	return DB.AutoMigrate(models.Models...)
}

// GetDB...
// Returns the database instance
func GetDB() *gorm.DB {
	return DB
}
