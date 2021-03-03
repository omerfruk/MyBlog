package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/routers"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	//db baglantıları ve migrate işlemi
	database.ConnectAndMigrate()
	routers.Router(app)
	// create islemleri
	/*service.CreateEntry("../img/bgCover.jpg","WELCOME TO MY PAGE","IT'S NICE TO MEET YOU","KNOW ME")

	service.CreatePortfolio("../img/computer.jpg","Computer","what is a computer and where are we in this")
	service.CreatePortfolio("../img/network.jpg","Network","How they get us all around with a click")
	service.CreatePortfolio("../img/user.jpg","User","Well, who are we users ")
	service.CreatePortfolio("../img/developer.jpg","Developer","what do developers do?")
	service.CreatePortfolio("../img/hardware.jpg","Hardware","what's the hardware? who are hardwareist")
	service.CreatePortfolio("../img/securty.jpg","Securty","who constitute our security?")

	service.CreateFooter("../img/fotom.jpg","ÖMER FARUK TASDEMIR","Developer","https://www.instagram.com/omer_fruk/?hl=tr","https://www.facebook.com/omerrf/","https://github.com/omerfruk","https://twitter.com/home?lang=tr")
	*/



	log.Fatal(app.Listen(":4747"))

	//app.Get("/info", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("info", fiber.Map{
	//		"Giris":       Omer.Title,
	//		"Name":        Omer.User,
	//		"Information": Omer.Information,
	//	})
	//})
	//app.Get("/Researches-Computer", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Computer!",
	//		"explanation":  Computer.Title,
	//		"explanation2": Computer2.Title,
	//		"explanation3": Computer3.Title,
	//		"Text":         Computer.Text,
	//		"Text2":        Computer2.Text,
	//		"Text3":        Computer3.Text,
	//	})
	//})
	//
	//app.Get("/Researches-User", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, User!",
	//		"explanation":  User.Title,
	//		"explanation2": User2.Title,
	//		"explanation3": User3.Title,
	//		"Text":         User.Text,
	//		"Text2":        User2.Text,
	//		"Text3":        User3.Text,
	//	})
	//})
	//app.Get("/Researches-Developer", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Developer!",
	//		"explanation":  Developer.Title,
	//		"explanation2": Developer2.Title,
	//		"explanation3": Developer3.Title,
	//		"Text":         Developer.Text,
	//		"Text2":        Developer2.Text,
	//		"Text3":        Developer3.Text,
	//	})
	//})
	//app.Get("/Researches-Hardware", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Hardware!",
	//		"explanation":  Hardware.Title,
	//		"explanation2": Hardware2.Title,
	//		"explanation3": Hardware3.Title,
	//		"Text":         Hardware.Text,
	//		"Text2":        Hardware2.Text,
	//		"Text3":        Hardware3.Text,
	//	})
	//})
	//app.Get("/Researches-Security", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Security!",
	//		"explanation":  Security.Title,
	//		"explanation2": Security2.Title,
	//		"explanation3": Security3.Title,
	//		"Text":         Security.Text,
	//		"Text2":        Security2.Text,
	//		"Text3":        Security3.Text,
	//	})
	//})
}



