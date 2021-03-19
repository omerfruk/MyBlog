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
		Fileurl := fmt.Sprintf("../img/kullanicilar/%s", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./img/kullanicilar/%s", file.Filename))
		service.SingUPUser(request.Information, request.Fullname, request.Mail, request.Password, Fileurl)
	} else {
		fmt.Println(err)
		service.SingUPUser(request.Information, request.Fullname, request.Mail, request.Password, "")
	}
	return c.Redirect("/login")
}
