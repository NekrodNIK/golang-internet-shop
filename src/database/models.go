package database

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UUID        string
	Name        string
	Description string
	Price       int
}
