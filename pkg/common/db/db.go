package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mongo_go/pkg/common/models"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Category{}, &models.Book{})

	return db
}
