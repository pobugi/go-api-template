package books

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Author      string `gorm:"type:varchar(255);not null"`
	Publication string `gorm:"type:varchar(255);not null"`
}

func (Book) TableName() string { return "books" }
