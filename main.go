package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
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
			"Info": "Hello, Info!",
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