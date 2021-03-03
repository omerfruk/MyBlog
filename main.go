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

	database.ConnectAndMigrate()

	routers.Router(app)
	//service.CreateTopBar("DeneMe","Anasayfa","Bilgilendirme","Arastirmalar","Iletisim")
	//service.CreateEntry("../img/bgCover.jpg","Hoşgeldiniz","Tanıştıgımıza memnun olduk","hakkımda")
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



