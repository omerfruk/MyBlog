package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/handlers"
)

func Router(app *fiber.App) {

	app.Static("/", "./")
	app.Get("/", handlers.IndexRender)
	app.Get("/info", handlers.InfoRender)
	app.Get("/login", handlers.Login)
	app.Post("/login", handlers.LogControl)
	app.Get("/logout", handlers.Logout)
	app.Get("/singup", handlers.SingUp)
	app.Post("/singup", handlers.SingUpPost)
	app.Get("/admin", handlers.AdminPage)
	app.Get("/down", handlers.DownPage)
	app.Post("/delete/:key", handlers.DeletUser)
	app.Post("/edit/:key", handlers.UpdateUser)
	app.Post("/create", handlers.CreateUser)
	app.Get("/gallery", handlers.Gallery)
	app.Get("/galleryAdd", handlers.GalleryAddGet)
	app.Post("/galleryAdd", handlers.GalleryAddPost)
	app.Get("/comments", handlers.Comments)
	app.Get("/comment", handlers.Comment)
	app.Post("/comment", handlers.CommentCreate)
	app.Get("/user/:key", handlers.GetUser)
	app.Post("/user/update/:key", handlers.UpdateUser)
	app.Get("/user/delete/:key", handlers.DeletUser)
	app.Get("/users", handlers.GetUsers)

	res := app.Group("/researcher/:key", handlers.ResearchersAll) // /researcher
	res.Get("/Computer", handlers.ResearchersAll)                 // /researcher/Computer
	res.Get("/User", handlers.ResearchersAll)                     // /researcher/User
	res.Get("/Network", handlers.ResearchersAll)                  // /researcher/Network
	res.Get("/Developer", handlers.ResearchersAll)                // /researcher/Developer
	res.Get("/Hardware", handlers.ResearchersAll)                 // /researcher/Hardware
	res.Get("/Security", handlers.ResearchersAll)                 // /researcher/Security

}
