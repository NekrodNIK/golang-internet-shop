package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbHandler struct {
	DbUrl      string
	connection *gorm.DB
}

func (DB DbHandler) connect() {
	con, err := gorm.Open(
		sqlite.Open(DB.DbUrl),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	DB.connection = con

	DB.connection.AutoMigrate(&Product{})
}
