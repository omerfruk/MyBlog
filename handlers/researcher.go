package handlers

import (
	"github.com/gofiber/fiber"
)

func ResearchersAll(c *fiber.Ctx) error {
	return c.Render("index.html",fiber.Map{

	})
}

