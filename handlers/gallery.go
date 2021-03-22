package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

func GalleryAddPost(c *fiber.Ctx) error {
	var request models.FotoGallery
	c.BodyParser(&request)
	file, err := c.FormFile("document")
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/gallery")
	}

	request.ImgSrc = fmt.Sprintf("../img/gallery/%s", file.Filename)
	c.SaveFile(file, fmt.Sprintf("./img/gallery/%s", file.Filename))
	service.CreateFoto(request)
	return c.Redirect("/gallery")

}

func GalleryAddGet(c *fiber.Ctx) error {
	return c.Render("galleryAdd", true)
}
func Gallery(c *fiber.Ctx) error {
	tempPhoto := service.GetFotoAll()
	tempFooter := service.GetFooter("OMER FARUK TASDEMIR")
	f := models.FotoStruct{
		FotoGallery: tempPhoto,
		FooterBar:   tempFooter,
	}
	return c.Render("gallery", f)

}
