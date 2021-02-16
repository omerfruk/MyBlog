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

	User:=Research{
		Model: gorm.Model{},
		Area:  "User",
		Title: "User",
		Text:  "User ile ilgili bilgi",
	}
	User2:=Research{
		Model: gorm.Model{},
		Area:  "User",
		Title: "User2",
		Text:  "User2 ile ilgili bilgi",
	}
	User3:=Research{
		Model: gorm.Model{},
		Area:  "User",
		Title: "User3",
		Text:  "User3 ile ilgili bilgi",
	}

	Developer:=Research{
		Model: gorm.Model{},
		Area:  "Developer",
		Title: "Developer",
		Text:  "Developer ile ilgili bilgi",
	}
	Developer2:=Research{
		Model: gorm.Model{},
		Area:  "Developer",
		Title: "Developer2",
		Text:  "Developer2 ile ilgili bilgi",
	}
	Developer3:=Research{
		Model: gorm.Model{},
		Area:  "Developer",
		Title: "Developer3",
		Text:  "Developer3 ile ilgili bilgi",
	}

	Hardware:=Research{
		Model: gorm.Model{},
		Area:  "Hardware",
		Title: "Hardware",
		Text:  "Hardware ile ilgili bilgi",
	}
	Hardware2:=Research{
		Model: gorm.Model{},
		Area:  "Hardware",
		Title: "Hardware2",
		Text:  "Hardware2 ile ilgili bilgi",
	}
	Hardware3:=Research{
		Model: gorm.Model{},
		Area:  "Hardware",
		Title: "Hardware3",
		Text:  "Hardware3 ile ilgili bilgi",
	}

	Security:=Research{
		Model: gorm.Model{},
		Area:  "Security",
		Title: "Security",
		Text:  "Security ile ilgili bilgi",
	}
	Security2:=Research{
		Model: gorm.Model{},
		Area:  "Security",
		Title: "Security2",
		Text:  "Security2 ile ilgili bilgi",
	}
	Security3:=Research{
		Model: gorm.Model{},
		Area:  "Security",
		Title: "Security3",
		Text:  "Security3 ile ilgili bilgi",
	}

	// verileri olusturma
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
/********************************************************************************************************/
	if networkTemp :=db.Where("title=?",Network.Title,"text=?",Network.Text).First(&Network);networkTemp.Error!=nil{
		db.Create(&Network)
	}
	if networkTemp :=db.Where("title=?",Network2.Title,"text=?",Network2.Text).First(&Network2);networkTemp.Error!=nil{
		db.Create(&Network2)
	}
	if networkTemp :=db.Where("title=?",Network3.Title,"text=?",Network3.Text).First(&Network3);networkTemp.Error!=nil{
		db.Create(&Network3)
	}
	/********************************************************************************************************/
	if userTepm:=db.Where("title=?",User.Title,"text=?",User.Text).First(&User);userTepm.Error!=nil{
		db.Create(&User)
	}
	if userTepm:=db.Where("title=?",User2.Title,"text=?",User2.Text).First(&User2);userTepm.Error!=nil{
		db.Create(&User2)
	}
	if userTepm:=db.Where("title=?",User3.Title,"text=?",User3.Text).First(&User3);userTepm.Error!=nil{
		db.Create(&User3)
	}
	/********************************************************************************************************/
	if developerTepm:=db.Where("title=?",Developer.Title,"text=?",Developer.Text).First(&Developer);developerTepm.Error!=nil{
		db.Create(&Developer)
	}
	if developerTepm:=db.Where("title=?",Developer2.Title,"text=?",Developer2.Text).First(&Developer2);developerTepm.Error!=nil{
		db.Create(&Developer2)
	}
	if developerTepm:=db.Where("title=?",Developer3.Title,"text=?",Developer3.Text).First(&Developer3);developerTepm.Error!=nil{
		db.Create(&Developer3)
	}
	/********************************************************************************************************/
	if hardwareTemp:=db.Where("title=?",Hardware.Title,"text=?",Hardware.Text).First(&Hardware);hardwareTemp.Error!=nil{
		db.Create(&Developer)
	}
	if hardwareTemp:=db.Where("title=?",Hardware2.Title,"text=?",Hardware2.Text).First(&Hardware2);hardwareTemp.Error!=nil{
		db.Create(&Hardware2)
	}
	if hardwareTemp:=db.Where("title=?",Hardware3.Title,"text=?",Hardware3.Text).First(&Hardware3);hardwareTemp.Error!=nil{
		db.Create(&Hardware3)
	}
	/********************************************************************************************************/
	if securityTemp:=db.Where("title=?",Security.Title,"text=?",Security.Text).First(&Security);securityTemp.Error!=nil{
		db.Create(&Security)
	}
	if securityTemp:=db.Where("title=?",Security2.Title,"text=?",Security2.Text).First(&Security2);securityTemp.Error!=nil{
		db.Create(&Security2)
	}
	if securityTemp:=db.Where("title=?",Security3.Title,"text=?",Security3.Text).First(&Security3);securityTemp.Error!=nil{
		db.Create(&Security3)
	}
	/********************************************************************************************************/

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
			"Text": Computer.Text,
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
			"Text": Network.Text,
			"Text2":Network2.Text,
			"Text3":Network3.Text,
		})
	})
	app.Get("/Researches-User", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, User!",
			"explanation":  User.Title,
			"explanation2": User2.Title,
			"explanation3": User3.Title,
			"Text": User.Text,
			"Text2":User2.Text,
			"Text3":User3.Text,
		})
	})
	app.Get("/Researches-Developer", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Developer!",
			"explanation":  Developer.Title,
			"explanation2": Developer2.Title,
			"explanation3": Developer3.Title,
			"Text": Developer.Text,
			"Text2":Developer2.Text,
			"Text3":Developer3.Text,
		})
	})
	app.Get("/Researches-Hardware", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Hardware!",
			"explanation":  Hardware.Title,
			"explanation2": Hardware2.Title,
			"explanation3": Hardware3.Title,
			"Text": Hardware.Text,
			"Text2":Hardware2.Text,
			"Text3":Hardware3.Text,
		})
	})
	app.Get("/Researches-Security", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Researches", fiber.Map{
			"Info":  "Hello, Security!",
			"explanation":  Security.Title,
			"explanation2": Security2.Title,
			"explanation3": Security3.Title,
			"Text": Security.Text,
			"Text2":Security2.Text,
			"Text3":Security3.Text,
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
