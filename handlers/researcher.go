package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ResearchersAll(c *fiber.Ctx) error {
	fmt.Println("geldi")
	return c.Render("researches", fiber.Map{
		"title": "homapages",
	})

}
func IndexRender(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title": "homapages",
	})
}
func InfoRender(c *fiber.Ctx) error {
	return c.Render("info", fiber.Map{
		"Giris":       "Omer.Title",
		"Name":        " Omer.User",
		"Information": "Omer.Information",
	})
}
