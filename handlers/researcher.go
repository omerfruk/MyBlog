package handlers

import (
	"github.com/gofiber/fiber"
)

func ResearchersAll(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Area" : "Area",
		"Title":"Title",
		"Text":"Text",
	})
}

