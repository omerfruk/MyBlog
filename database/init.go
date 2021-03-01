package database

import (
	"fmt"

	"github.com/omerfruk/my-blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	DATABASE = "blog"
	USER     = "postgres"
	PASSWORD = "123"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Connect() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := gorm.Open(postgres.Open(vt), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
}

func ConnectAndMigrate() {
	Connect()
	Migrate()
}

func Migrate() {
	// postgres şemaları burdaki olaylardan haberdar olacak
	db.AutoMigrate(&models.User{})
}

