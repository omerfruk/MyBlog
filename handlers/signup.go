package handlers

import (
	"fmt"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"

	"github.com/gofiber/fiber/v2"
)

//gingup render
func SingUp(c *fiber.Ctx) error {

	return c.Render("singup", true)
}

//sing up post render
func SingUpPost(c *fiber.Ctx) error {
	var request models.User
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	file, err := c.FormFile("document")
	if err == nil {
		request.ImgSrc = fmt.Sprintf("../assets/img/kullanicilar/%s", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./assets/img/kullanicilar/%s", file.Filename))

	} else {
		fmt.Println(err)
		request.ImgSrc = ""
	}
	service.CreateUser(request, false)
	return c.Redirect("/login")
}
