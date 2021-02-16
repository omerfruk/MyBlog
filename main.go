package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	DATABASE = "blog"
	USER     = "postgres"
	PASSWORD = "123"
)
type Admin struct {
	gorm.Model
	Title string
	User string
	Information string
}
type Research struct {
	gorm.Model
	Area string
	Title string
	Text string
}
func main() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := gorm.Open(postgres.Open(vt), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// postgres şemaları burdaki olaylardan haberdar olacak
	db.AutoMigrate(&Admin{})

	Omer := Admin{
		Model:       gorm.Model{},
		Title:       "Merhaba",
		User:        "Ben Ömer Faruk",
		Information: "İzmirin konak ilçesinde ",
	}

	// eger db de Omer.User adlı bir veri eşleşmesi varsa onu tekrardan oluşturma
	if result := db.Where("User=?", Omer.User); result.Error != nil {
		db.Create(&Omer)
	}
	Computer:=Research{
		Model: gorm.Model{},
		Area: "computer info",
		Title: "bilgisayar 1",
		Text:  "bilgisayarla alakalı veri 1",
	}
	Computer2:=Research{
		Model: gorm.Model{},
		Area: "computer info",
		Title: "bilgisayar 2",
		Text:  "bilgisayarla alakalı veri 2",
	}
	Computer3:=Research{
		Model: gorm.Model{},
		Area: "computer info",
		Title: "bilgisayar 3",
		Text:  "bilgisayarla alakalı veri 3",
	}
	Network:=Research{
		Model: gorm.Model{},
		Area:  "Network",
		Title: "Network",
		Text:  "Network ile ilgili bilgi",
	}
	Network2:=Research{
		Model: gorm.Model{},
		Area:  "Network",
		Title: "Network2",
		Text:  "Network2 ile ilgili bilgi",
	}
	Network3:=Research{
		Model: gorm.Model{},
		Area:  "Network",
		Title: "Network3",
		Text:  "Network3 ile ilgili bilgi",
	}
	db.AutoMigrate(&Research{})
	if computerTemp := db.Where("Title=?",Computer.Title,"Text=?",Computer.Text).First(&Computer);computerTemp.Error!=nil{
		db.Create(&Computer)
	}
	if computerTemp := db.Where("Title=?",Computer2.Title,"Text=?",Computer2.Text).First(&Computer2);computerTemp.Error!=nil{
		db.Create(&Computer2)
	}
	if computerTemp := db.Where("Title=?",Computer3.Title,"Text=?",Computer3.Text).First(&Computer3);computerTemp.Error!=nil{
		db.Create(&Computer3)
	}
	if networkTemp :=db.Where("title=?",Network.Title,"text=?",Network.Text).First(&Network);networkTemp.Error!=nil{
		db.Create(&Network)
	}
	if networkTemp :=db.Where("title=?",Network2.Title,"text=?",Network2.Text).First(&Network2);networkTemp.Error!=nil{
		db.Create(&Network2)
	}
	if networkTemp :=db.Where("title=?",Network3.Title,"text=?",Network3.Text).First(&Network3);networkTemp.Error!=nil{
		db.Create(&Network3)
	}

	// Initialize standard Go html template engine
	engine := html.New("./", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Main": "Hello, main!",
		})
	})
	app.Get("/info", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("info", fiber.Map{
			"Giris":       Omer.Title,
			"Name":        Omer.User,
			"Information": Omer.Information,
		})
	})
	app.Get("/Researches-Computer", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Computer!",
			"explanation":  Computer.Title,
			"explanation2": Computer2.Title,
			"explanation3": Computer3.Title,
			"Text":Computer.Text,
			"Text2":Computer2.Text,
			"Text3":Computer3.Text,
		})
	})
	app.Get("/Researches-Network", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Network!",
			"explanation":  Network.Title,
			"explanation2": Network2.Title,
			"explanation3": Network3.Title,
			"Text":Network.Text,
			"Text2":Network2.Text,
			"Text3":Network3.Text,
		})
	})
	app.Get("/Researches-User", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, User!",
			"dene":  "Network1",
			"dene2": "Network2",
			"dene3": "Network3",
		})
	})
	app.Get("/Researches-Developer", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Developer!",
			"dene":  "Network1",
			"dene2": "Network2",
			"dene3": "Network3",
		})
	})
	app.Get("/Researches-Hardware", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Hardware!",
			"dene":  "Network1",
			"dene2": "Network2",
			"dene3": "Network3",
		})
	})
	app.Get("/Researches-Security", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Security!",
			"dene":  "Network1",
			"dene2": "Network2",
			"dene3": "Network3",
		})
	})

	//app.Get("/deneme", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("deneme", fiber.Map{
	//		"Title": "Hello, World!",
	//	})
	//})

	app.Listen(":4747")
}
