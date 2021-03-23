package routers

import (
	"fmt"
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
	app.Get("/down", handlers.DownPage)

	app.Get("/gallery", handlers.Gallery)

	app.Get("/comments", handlers.Comments)

	app.Get("/users", handlers.GetUsers)

	res := app.Group("/researcher/:key", handlers.ResearchersAll) // /researcher
	res.Get("/Computer", handlers.ResearchersAll)                 // /researcher/Computer
	res.Get("/User", handlers.ResearchersAll)                     // /researcher/User
	res.Get("/Network", handlers.ResearchersAll)                  // /researcher/Network
	res.Get("/Developer", handlers.ResearchersAll)                // /researcher/Developer
	res.Get("/Hardware", handlers.ResearchersAll)                 // /researcher/Hardware
	res.Get("/Security", handlers.ResearchersAll)                 // /researcher/Security

	adm := app.Group("/admin", sessionControl)
	adm.Get("/", handlers.AdminPage)
	adm.Get("/comment", handlers.Comment)
	adm.Post("/delete/:key", handlers.DeletUser)
	adm.Post("/edit/:key", handlers.UpdateUser)
	adm.Post("/create", handlers.CreateUser)
	adm.Get("/galleryAdd", handlers.GalleryAddGet)
	adm.Post("/galleryAdd", handlers.GalleryAddPost)
	adm.Get("/user/:key", handlers.GetUser)
	adm.Post("/user/update/:key", handlers.UpdateUser)
	adm.Get("/user/delete/:key", handlers.DeletUser)

}

func sessionControl(c *fiber.Ctx) error {
	sess, err := handlers.SessionStore.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	if sess.Get("temp") == nil {
		return c.Redirect("/login")
	}
	return c.Next()

}
