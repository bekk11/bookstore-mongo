package models

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Image       string    `gorm:"null" json:"image"`
	Title       string    `gorm:"not null" json:"title"`
	Author      string    `gorm:"not null" json:"author"`
	Description string    `gorm:"not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	PubDate     time.Time `gorm:"null" json:"pubDate"`
	CategoryID  uint      `gorm:"not null" json:"categoryID"`
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
