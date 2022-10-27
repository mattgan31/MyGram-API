package database

import (
	"final-project-fga/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "alam"
	dbport   = 5432
	dbname   = "final_fga"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.Photo{})
	db.Debug().AutoMigrate(models.SocialMedia{})
	db.Debug().AutoMigrate(models.Comment{})

}

func GetDB() *gorm.DB {
	return db
}
