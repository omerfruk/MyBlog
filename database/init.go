package database

import (
	"fmt"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	var err error
	db, err = gorm.Open(postgres.Open(vt), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
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
	db.AutoMigrate(&models.Topbar{})
	db.AutoMigrate(&models.Instructions{})
	db.AutoMigrate(&models.FooterBar{})
	db.AutoMigrate(&models.Portfolio{})
	db.AutoMigrate(&models.Entry{})
	db.AutoMigrate(&models.Research{})

}
