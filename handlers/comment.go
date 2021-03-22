package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
)

//dusunceler sayfasi
func Comments(c *fiber.Ctx) error {
	temp := service.GetAllComments()
	return c.Render("comments", temp)
}

//dusunce sayfasi get
func Comment(c *fiber.Ctx) error {
	return c.Render("comment", true)
}

//dusunce sayfasi olusturmna
func CommentCreate(c *fiber.Ctx) error {
	var temp models.Comment
	err := c.BodyParser(&temp)
	if err != nil {
		fmt.Println(err)
	}
	service.CommentCreate(temp)
	return c.Redirect("/comment")
}
