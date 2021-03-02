package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/service"
)

func ResearchersAll(c *fiber.Ctx) error {
	return c.Render("researcher", fiber.Map{
		"Info": "Hello, " + c.Params("key") + "!",
		//		"explanation":  ,
		//		"explanation2": ,
		//		"explanation3": ,
		//		"Text":         ,
		//		"Text2":        ,
		//		"Text3":        ,
	})

}
func IndexRender(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title": "homapages",
	})
}
func InfoRender(c *fiber.Ctx) error {
	user := service.GetUser("Ömer Faruk")
	return c.Render("info", fiber.Map{
		"Giris":   	   user.Header,
		"Name":        user.Fullname,
		"Information": user.Information,
	})
}
