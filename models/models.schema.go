package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    string
	Name  string
	Email string
}
type Post struct {
	gorm.Model
	ID      string
	UserId  string
	Content string
}

func CreateTables(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
