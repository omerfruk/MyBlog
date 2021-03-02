package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
	"io/ioutil"
	"os"
)

func ResearchersAll(c *fiber.Ctx) error {
	var research []models.Research
	jsonFile, err := os.Open("fakedata/fake_data.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byteValue, &research)
	if err != nil {
		fmt.Println(err)
	}
	/*var titles []string
	var text []string
	j:=0*/
	for i := 0; i < len(research); i++ {
		if research[i].Area == c.Params("key") {
			fmt.Println(research[i].Area)
		/*	titles[j] = research[i].Title
			fmt.Println(titles[i])
			text[j] = research[i].Text
			fmt.Println(text[i])
			j++*/
		}
	}
	return c.Render("researcher", fiber.Map{
		"Info":         "Hello, " + c.Params("key") + "!",
		/*"explanation":  titles[0],
		"explanation2": titles[1],
		"explanation3": titles[2],
		"Text":         text[0],
		"Text2":        text[1],
		"Text3":        text[2],*/
	})

}

/*func Unmurshal(key string) *models.Research {

	var need models.Research
	for i:=0;i<len(research);i++ {
		if key == research[i].Title{
			need = research[i]
			fmt.Println(need)
			return &need
		}
	}
	return &need
}*/
func IndexRender(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title": "homapages",
	})
}
func InfoRender(c *fiber.Ctx) error {
	user := service.GetUser("Ömer Faruk")
	return c.Render("info", fiber.Map{
		"ID":          user.ID,
		"Giris":       user.Header,
		"Name":        user.Fullname,
		"Information": user.Information,
		"Authority":   user.Authority,
	})
}
