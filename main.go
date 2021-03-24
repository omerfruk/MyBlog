package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/routers"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	//db baglantıları ve migrate işlemi
	database.ConnectAndMigrate()
	routers.Router(app)

	log.Fatal(app.Listen(":4747"))

}
