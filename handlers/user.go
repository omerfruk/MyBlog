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
	temp := service.GetUserId(c.Params("key"))
	return c.Render("user", temp)
}

func Getadmin(c *fiber.Ctx) error {
	temp := service.GetAdmin("omer faruk")
	return c.Render("user", temp)
}

//admin sayufasidan  kullanici olusturmak icin
func CreateUser(c *fiber.Ctx) error {
	var temp models.User
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	service.CreateUser(temp, false)
	return c.Redirect("/admin")
}

//admin sayfasindan kullanicilari editlemek icin
func UpdateUser(c *fiber.Ctx) error {
	key := c.Params("key")
	var request, temp models.User
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	database.DB().Where("id = ?", key).First(&temp)
	request.ID = temp.ID
	file, err := c.FormFile("document")

	url := fmt.Sprintf("/admin/user/%d", temp.ID)
	if request.Mail == "" && request.Fullname == "" {
		fmt.Println("degerler bos olamaz")
		return c.Redirect(url)
	}
	if err == nil {
		request.ImgSrc = fmt.Sprintf("../img/kullanicilar/%s", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./img/kullanicilar/%s", file.Filename))

	} else {
		request.ImgSrc = temp.ImgSrc
		fmt.Println("fotograf bulunmadigi icin eski fotografla kullanilacak")
	}
	service.UpdateUser(request, temp)
	return c.Redirect(url)
}

//admin sayfasindan kullanici silmek icin
func DeletUser(c *fiber.Ctx) error {
	key := c.Params("key")
	service.DeleteUser(key)
	return c.Redirect("/admin")
}
