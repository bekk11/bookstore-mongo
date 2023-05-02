package main

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}
