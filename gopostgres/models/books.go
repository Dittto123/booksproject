package models

import "gorm.io/gorm"

type Books struct {
	ID        int     `gorm:"primary key; autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Books{})
	return err
}
