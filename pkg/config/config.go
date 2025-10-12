package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTPAddr     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
	DBLoc  string
}

func Load() AppConfig {
	v := viper.New()

	// Defaults (use native types for durations)
	v.SetDefault("http_addr", ":9010")
	v.SetDefault("read_timeout", 15*time.Second)
	v.SetDefault("write_timeout", 15*time.Second)

	v.SetDefault("db_host", "127.0.0.1")
	v.SetDefault("db_port", 3306)
	v.SetDefault("db_user", "myuser")
	v.SetDefault("db_pass", "mypassword")
	v.SetDefault("db_name", "bookstore")
	v.SetDefault("db_loc", "Local")

	// Optional .env file
	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		if err := v.ReadInConfig(); err != nil {
			log.Printf("warn: .env read failed: %v", err)
		}
	}

	// Environment variables override everything
	// e.g. HTTP_ADDR, READ_TIMEOUT, DB_HOST, DB_PORT, ...
	v.AutomaticEnv()
	// IMPORTANT: never pass nil here (viper 1.21 would panic)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // allow dotted keys if you add them later

	var cfg AppConfig
	// Instead of Unmarshal (which can be finicky with durations) read fields explicitly:
	cfg.HTTPAddr = v.GetString("http_addr")
	cfg.ReadTimeout = v.GetDuration("read_timeout")
	cfg.WriteTimeout = v.GetDuration("write_timeout")

	cfg.DBHost = v.GetString("db_host")
	cfg.DBPort = v.GetInt("db_port")
	cfg.DBUser = v.GetString("db_user")
	cfg.DBPass = v.GetString("db_pass")
	cfg.DBName = v.GetString("db_name")
	cfg.DBLoc = v.GetString("db_loc")

	return cfg
}

func (c AppConfig) MySQLDSN() string {
	// parseTime=true is important for GORM time fields
	loc := c.DBLoc
	if loc == "" {
		loc = "Local"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, loc)
}

//var db *gorm.DB
//
//func Connect() {
//	dsn := "myuser:mypassword@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
//
//	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to the database: %v", err)
//	}
//	db = d
//}
//
//func GetDB() *gorm.DB {
//	return db
//}
