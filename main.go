package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/routers"
	"github.com/omerfruk/my-blog/service"
	"log"
)

func main() {
	engine := html.New("views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	//db baglantıları ve migrate işlemi
	database.ConnectAndMigrate()
	routers.Router(app)
	// create islemleri
	/*service.CreateEntry("../img/bgCover.jpg","WELCOME TO MY PAGE","IT'S NICE TO MEET YOU","KNOW ME")

	service.CreatePortfolio("../img/computer.jpg","Computer","what is a computer and where are we in this")
	service.CreatePortfolio("../img/network.jpg","Network","How they get us all around with a click")
	service.CreatePortfolio("../img/user.jpg","User","Well, who are we users ")
	service.CreatePortfolio("../img/developer.jpg","Developer","what do developers do?")
	service.CreatePortfolio("../img/hardware.jpg","Hardware","what's the hardware? who are hardwareist")
	service.CreatePortfolio("../img/securty.jpg","Securty","who constitute our security?")

	service.CreateFooter("../img/fotom.jpg","ÖMER FARUK TASDEMIR","Developer","https://www.instagram.com/omer_fruk/?hl=tr","https://www.facebook.com/omerrf/","https://github.com/omerfruk","https://twitter.com/home?lang=tr")

	service.CreateInstructions("Let's learn something about technology","The most efficient thing which is invented by people is book","Computer","Web","Users","It is certain that the development progress of the computer without pausing, from the past to the present,\nwill continue exponentially. So where are we in this technology?","We can go from one side of the world to the other with one click. \nSo how do we make this journey?","We, I mean the users, can shape the progress of technology. It sounds amazing right?")

	service.CreateTopBar("ÖmFar.","home","future","researcher","contact","Login")


	*/
	//service.CreateUser("../img/fotom.jpg","Merhabalar","Ömer Faruk","Izmir dogumlu bir insanım",true)
	service.CreatePortfolio("../img/game.jpg","Game","Of course, the game everyone wants!")
		log.Fatal(app.Listen(":4747"))


}



