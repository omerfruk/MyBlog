package main

import (
	"fmt"
	"MyBlog/structure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"


)
const (
	HOST = "localhost"
	DATABASE = "blog"
	USER = "postgres"
	PASSWORD = "123"
)

func main() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := gorm.Open(postgres.Open(vt), &gorm.Config{})
	if err !=nil{
		fmt.Println(err)
	}
	// postgres şemaları burdaki olaylardan haberdar olacak
	db.AutoMigrate(&struct.Admin{})
	Omer := Admin {
		Model:       gorm.Model{},
		Title:       "Merhaba",
		User:        "Ömer Faruk",
		Information: "İzmirin konak ilçesinde ",
	}
	// eger db de Omer.User adlı bir veri eşleşmesi varsa onu tekrardan oluşturma
	if result :=db.Where("User=?",Omer.User);result.Error!=nil{
		db.Create(&Omer)
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
			"Giris":Omer.Title,
			"Name": Omer.User,
			"Information":Omer.Information,
		})
	})
	app.Get("/Researches-Computer", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, Computer!",
			"dene":"computer1",
			"dene2":"computer2",
			"dene3":"computer3",
		})
	})
	app.Get("/Researches-Network", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, Network!",
			"dene":"Network1",
			"dene2":"Network2",
			"dene3":"Network3",
		})
	})
	app.Get("/Researches-User", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, User!",
			"dene":"Network1",
			"dene2":"Network2",
			"dene3":"Network3",
		})
	})
	app.Get("/Researches-Developer", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, Developer!",
			"dene":"Network1",
			"dene2":"Network2",
			"dene3":"Network3",
		})
	})
	app.Get("/Researches-Hardware", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, Hardware!",
			"dene":"Network1",
			"dene2":"Network2",
			"dene3":"Network3",
		})
	})
	app.Get("/Researches-Security", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info": "Hello, Security!",
			"dene":"Network1",
			"dene2":"Network2",
			"dene3":"Network3",
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