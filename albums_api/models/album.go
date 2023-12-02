package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	ID     string  `json:"id" gorm:"primary_key"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}
