package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

func GetUsers(c *fiber.Ctx) error {
	var user []models.User
	err := database.DB().Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return c.Render("users", user)
}

//user sayfasi get
func GetUser(c *fiber.Ctx) error {
	key := c.Params("key")
	var temp models.User
	database.DB().Where("id = ?", key).Find(&temp)
	return c.Render("user", temp)
}

//admin sayfasindan kullanicilari editlemek icin
func EditUser(c *fiber.Ctx) error {
	key := c.Params("key")
	var request, temp models.User
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	database.DB().Where("id = ?", key).First(&temp)

	file, err := c.FormFile("document")

	url := fmt.Sprintf("/user/%d", temp.ID)
	if request.Mail == "" && request.Fullname == "" {
		fmt.Println("degerler bos olamaz")
		return c.Redirect(url)
	}
	if err == nil {
		Fileurl := fmt.Sprintf("../img/kullanicilar/%s", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./img/kullanicilar/%s", file.Filename))
		service.UpdateUser(temp.ID, request.Mail, request.Fullname, request.Information, Fileurl)

	} else {
		fmt.Println(err)
		service.UpdateUser(temp.ID, request.Mail, request.Fullname, request.Information, temp.ImgSrc)
	}
	return c.Redirect(url)
}

//admin sayfasindan kullanici silmek icin
func DltUSer(c *fiber.Ctx) error {
	key := c.Params("key")

	service.DeleteUser(key)

	return c.Redirect("/admin")
}

//admin sayufasidan  kullanici olusturmak icin
func CreateUser(c *fiber.Ctx) error {
	var temp models.RequestSingUp
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	service.SingUPUser(temp.Phone, temp.FullName, temp.Email, "", "")
	return c.Redirect("/admin")
}
