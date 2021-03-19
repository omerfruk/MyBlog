package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

//login sayfasının renderi
func Login(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	//eger session acık ise girilecek yer
	if sess.Get("temp") != nil {
		// Destroy session
		if err := sess.Destroy(); err != nil {
			panic(err)
		}
		return c.Redirect("/")
	} else {
		//sessions açık değilse girilecek  method
		return c.Render("login", true)
	}
}

//login kontrol
func LogControl(c *fiber.Ctx) error {
	var request models.RequestBody
	var temp models.User

	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	pass := service.Sha256String(request.Password)
	err = database.DB().Where("mail = ? and password = ?", request.Email, pass).First(&temp).Error
	if err != nil {
		return c.Redirect("down")
	}

	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}

	defer sess.Save()

	sess.Set("temp", temp.Fullname)
	fmt.Println(sess.Get("temp"))

	return c.Redirect("/admin")
}

//yanlıs giris sayaci
var sayac int = 0

//yanlıs giris ekrani
func DownPage(c *fiber.Ctx) error {
	sayac++
	if sayac >= 3 {
		sayac = 0
		c.Redirect("/singup")
	}
	conn := "Your Login Failed"
	return c.Render("down", conn)

}
