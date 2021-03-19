package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

func GalleryAddPost(c *fiber.Ctx) error {
	var request models.FotoGallery
	c.BodyParser(&request)
	file, err := c.FormFile("document")
	if err != nil {
		fmt.Println(err)
		c.Redirect("/gallery")
	} else {
		Fileurl := fmt.Sprintf("../img/gallery/%s", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./img/gallery/%s", file.Filename))
		service.CreateFoto(Fileurl, request.Title, request.Text)
		return c.Redirect("/gallery")
	}
	return c.Redirect("/gallery")
}

func GalleryAddGet(c *fiber.Ctx) error {
	return c.Render("galleryAdd", true)
}
func Gallery(c *fiber.Ctx) error {
	var tempPhoto []models.FotoGallery
	var tempFooter models.FooterBar
	database.DB().Find(&tempPhoto)
	database.DB().Find(&tempFooter)
	f := models.FotoStruct{
		FotoGallery: tempPhoto,
		FooterBar:   tempFooter,
	}
	return c.Render("gallery", f)

}
