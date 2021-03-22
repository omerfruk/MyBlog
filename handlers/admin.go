package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

//admin sayfasi
func AdminPage(c *fiber.Ctx) error {
	var admin models.AdminPage
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sess.Get("temp"))
	if sess.Get("temp") == nil {
		return c.Redirect("/login")
	}

	var user []models.User
	inf := service.GetUser("omer faruk")
	topbar := service.GetTopBar("ÖmFar.")
	err = database.DB().Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	admin = models.AdminPage{
		Topbar: topbar,
		Info:   inf,
		User:   user,
	}
	return c.Render("admin", admin)

}
