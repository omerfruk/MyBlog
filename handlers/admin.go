package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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
	if sess.Get("temp") == nil {
		return c.Redirect("/login")
	}

	inf := service.GetUser(sess.Get("mail"))
	topbar := service.GetTopBar("ÖmFar.")
	user, _ := service.Find()
	admin = models.AdminPage{
		Topbar: topbar,
		Info:   inf,
		User:   user,
	}
	return c.Render("admin", admin)

}
