package database

import (
	"fmt"
	"log"
	"movie_festival_app/src/models"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func LoadDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USERDB"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("DBPORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	DB.Debug().AutoMigrate(models.User{}, models.Movie{})
}

func GetDB() *gorm.DB {
	return DB
}
