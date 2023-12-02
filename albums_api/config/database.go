package config

import (
	"github.com/orket-sam/go_lessons/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Album{})
	DB = db
}
