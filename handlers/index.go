package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

//index sayfasnini renderlemek icin
func IndexRender(c *fiber.Ctx) error {

	topbar := service.GetTopBar("ÖmFar.")
	entry := service.GetEntry("WELCOME TO MY PAGE")
	intro := service.GetInstructions("Let's learn something about technology")
	footer := service.GetFooter("OMER FARUK TASDEMIR")
	portfolio := service.GetPortfolio()
	sess, err := SessionStore.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	a := models.Anasayfa{
		Portfolio: portfolio,
		Entry:     entry,
		Topbar:    topbar,
		Intro:     intro,
		Footer:    footer,
		Bool:      sess.Fresh(),
	}

	return c.Render("index", a)
}
