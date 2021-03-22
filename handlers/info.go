package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

//info Render bolumu
func InfoRender(c *fiber.Ctx) error {
	topbar := service.GetTopBar("ÖmFar.")
	user := service.GetAdmin("omer faruk")
	footer := service.GetFooter("OMER FARUK TASDEMIR")

	I := models.Info{
		Topbar: topbar,
		User:   user,
		Footer: footer,
	}
	return c.Render("info", I)
}
