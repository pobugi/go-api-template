package db

import (
	"log"

	"github.com/pobugi/go-crud-mysql/pkg/books"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/pobugi/go-crud-mysql/pkg/config"
)

func Open(cfg config.AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB open err: %v", err)
	}

	// Automigrate is fine for demo. in prod use sql migrations
	if err := db.AutoMigrate(&books.Book{}); err != nil {
		log.Fatalf("DB migrate err: %v", err)
	}
	return db

}
