package main

import (

	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)


func main() {

	database.ConnectAndMigrate()


	CreateUser("omer","dsadsa",true)
	//app.Static("/", "./")
	//app.Get("/", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("index", fiber.Map{
	//		"Main": "Hello, main!",
	//	})
	//})

	//app.Get("/info", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("info", fiber.Map{
	//		"Giris":       Omer.Title,
	//		"Name":        Omer.User,
	//		"Information": Omer.Information,
	//	})
	//})
	//app.Get("/Researches-Computer", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Computer!",
	//		"explanation":  Computer.Title,
	//		"explanation2": Computer2.Title,
	//		"explanation3": Computer3.Title,
	//		"Text":         Computer.Text,
	//		"Text2":        Computer2.Text,
	//		"Text3":        Computer3.Text,
	//	})
	//})
	//
	//app.Get("/Researches-User", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, User!",
	//		"explanation":  User.Title,
	//		"explanation2": User2.Title,
	//		"explanation3": User3.Title,
	//		"Text":         User.Text,
	//		"Text2":        User2.Text,
	//		"Text3":        User3.Text,
	//	})
	//})
	//app.Get("/Researches-Developer", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Developer!",
	//		"explanation":  Developer.Title,
	//		"explanation2": Developer2.Title,
	//		"explanation3": Developer3.Title,
	//		"Text":         Developer.Text,
	//		"Text2":        Developer2.Text,
	//		"Text3":        Developer3.Text,
	//	})
	//})
	//app.Get("/Researches-Hardware", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Hardware!",
	//		"explanation":  Hardware.Title,
	//		"explanation2": Hardware2.Title,
	//		"explanation3": Hardware3.Title,
	//		"Text":         Hardware.Text,
	//		"Text2":        Hardware2.Text,
	//		"Text3":        Hardware3.Text,
	//	})
	//})
	//app.Get("/Researches-Security", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("Researches", fiber.Map{
	//		"Info":         "Hello, Security!",
	//		"explanation":  Security.Title,
	//		"explanation2": Security2.Title,
	//		"explanation3": Security3.Title,
	//		"Text":         Security.Text,
	//		"Text2":        Security2.Text,
	//		"Text3":        Security3.Text,
	//	})
	//})

	//app.Get("/deneme", func(c *fiber.Ctx) error {
	//	// Render index template
	//	return c.Render("deneme", fiber.Map{
	//		"Title": "Hello, World!",
	//	})
	//})
}

func CreateResearch (area  string,title string,text  string){
	temp:=new(models.Research)
	temp.Area=area
	temp.Text=text
	temp.Title=title

	if researchTemp := database.DB().Where("title=? and text=?",temp.Title,temp.Text).First(&temp); researchTemp.Error!=nil{
		database.DB().Create(&temp)
	}
}
func CreateUser(name string,info string,isAdmin bool)  {
	temp:=new(models.User)
	temp.Fullname=name
	temp.Information=info
	if isAdmin == true{
		temp.Authority=models.Admin
	}else
	{
		temp.Authority=models.Basic
	}

	if UserTemp:=database.DB().Where("fullname=?",temp.Fullname).First(&temp); UserTemp.Error!=nil{
		database.DB().Create(&temp)
	}
}