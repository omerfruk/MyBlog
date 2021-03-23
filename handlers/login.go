package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
	"time"
)

// session olusturma
var store = session.New(session.Config{
	Expiration:   24 * time.Hour,
	CookieName:   "session_id",
	KeyGenerator: utils.UUID,

	//storage problemi olusuyor sonrasında bununla ilgilenilecek
	/*Storage: postgres.New(postgres.Config{
		Host:       "localhost",
		Port:       5432,
		Username:   "postgres",
		Password:   "123",
		Database:   "blog",
		Table:      "fiber_storage",
		GCInterval: 10 * time.Hour,
	}),*/
})

//login sayfasının renderi
func Login(c *fiber.Ctx) error {
	return c.Render("login", true)

}

//login kontrol
func LogControl(c *fiber.Ctx) error {
	var request models.RequestBody
	var temp models.User

	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	pass := service.Sha256String(request.Password)
	err = database.DB().Where("mail = ? and password = ?", request.Email, pass).First(&temp).Error
	if err != nil {
		return c.Redirect("down")
	}

	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}

	defer sess.Save()

	sess.Set("temp", temp.Fullname)
	sess.Set("mail", temp.Mail)

	return c.Redirect("/admin")
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	sess.Destroy()
	return c.Redirect("/")
}

//yanlıs giris sayaci
var sayac int = 0

//yanlıs giris ekrani
func DownPage(c *fiber.Ctx) error {
	sayac++
	if sayac >= 3 {
		sayac = 0
		c.Redirect("/singup")
	}
	conn := "Your Login Failed"
	return c.Render("down", conn)

}

// jwt uretimi
/*func LogControlJWT(c *fiber.Ctx) error {
	var request models.RequestBody
	var temp models.User
	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request.Email + "   " + request.Password)
	err = database.DB().Where("mail = ? and password = ?", request.Email, request.Password).First(&temp).Error
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	fmt.Println(temp.Fullname)
	fmt.Println(temp.Authority)
	token := jwt.New(jwt.SigningMethodHS256)

	claims:=token.Claims.(jwt.MapClaims)
	claims["name"]=temp.Fullname
	claims["admin"]=temp.Authority
	claims["exp"]=time.Now().Add(time.Hour * 24 * 30).Unix()

	fmt.Println(claims)

	t,err := token.SignedString([]byte("secret"))
	if err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"token":t,
	})
}
func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}*/
