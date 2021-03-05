package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/my-blog/models"
	"github.com/omerfruk/my-blog/service"
	"io/ioutil"
	"os"
	"strings"
)

//db den okuyup sayfaya yansıtmak için
func ResearchersAll(c *fiber.Ctx) error {
	var res  [] models.Research
	area := c.Params("key")
	temp:=strings.ToLower(area)
	res = service.GetResearch(temp)
	return c.Render("researcher",fiber.Map{
		"res": res,
	})
}

//json dosyasından okumak için
func ResearchersAllJson(c *fiber.Ctx) error {
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
	var titles [3]string
	var text [3]string
	j := 0
	area := c.Params("key")
	for i := 0; i < len(research); i++ {
		if research[i].Area == area {
			//fmt.Println(research[i].Area)
			titles[j] = research[i].Title
			//fmt.Println(titles[j])
			text[j] = research[i].Text
			//fmt.Println(text[j])
			j++
		}
	}
	return c.Render("researcher", fiber.Map{
		"Info":         "Hello, " + area + " !",
		"explanation":  titles[0],
		"explanation2": titles[1],
		"explanation3": titles[2],
		"Text":         text[0],
		"Text2":        text[1],
		"Text3":        text[2],
	})

}

/*func DenemeRender(c *fiber.Ctx)error{
	key :=c.Params("key")
	var res = service.GetResearch(key)

	return c.Render("deneme",res)
}*/

func IndexRender(c *fiber.Ctx) error {
	topbar := service.GetTopBar("ÖmFar.")
	entry := service.GetEntry("WELCOME TO MY PAGE")
	intro := service.GetInstructions("Let's learn something about technology")
	footer := service.GetFooter("OMER FARUK TASDEMIR")
	portfolio := service.GetPortfolio()
	type Anasayfa struct {
		Portfolio []models.Portfolio
		Research models.Research
	}
	a := Anasayfa{
		Portfolio: portfolio,
	}
	return c.Render("index", fiber.Map{

		"portfolio": portfolio,

		"title": "homapages",
		//topbar bolumu
		"logo":    topbar.Logo,
		"home":    topbar.Home,
		"future":  topbar.Future,
		"prtfl":   topbar.Port,
		"contact": topbar.Contact,
		//giris bolumu
		"firs_img":     entry.ImgSrc,
		"entry_header": entry.Header,
		"entry_text":   entry.Text,
		"entry_button": entry.ButtonText,
		//informations bolumu
		"headerIntro": intro.Title,
		"textIntro":   intro.Info,
		"LI":          intro.LeftIntro,
		"MI":          intro.MidIntro,
		"RI":          intro.RghtIntro,
		"LD":          intro.LeftDesc,
		"MD":          intro.MidDesc,
		"RD":          intro.RghtDesc,
		//footerBar bölümü
		"footer_img":        footer.ImgSrc,
		"footer_name":       footer.FullName,
		"footer_employment": footer.Employment,
		"footer_insta":      footer.InstaSrc,
		"footer_facc":       footer.FaccSrc,
		"footer_git":        footer.GitSrc,
		"footer_tw":         footer.TwSrec,
	})
}
func InfoRender(c *fiber.Ctx) error {
	topbar := service.GetTopBar("ÖmFar.")
	user := service.GetUser("Ömer Faruk")
	footer := service.GetFooter("OMER FARUK TASDEMIR")

	return c.Render("info", fiber.Map{
		//topBar bölümü
		"logo":      topbar.Logo,
		"home":      topbar.Home,
		"future":    topbar.Future,
		"portfolio": topbar.Port,
		"contact":   topbar.Contact,
		//usere bolumu
		"img":         user.ImgSrc,
		"Giris":       user.Header,
		"Name":        user.Fullname,
		"Information": user.Information,
		"Authority":   user.Authority,
		//footerBar bolumunden cekilenler
		"insta": footer.InstaSrc,
		"face":  footer.FaccSrc,
		"git":   footer.GitSrc,
		"tw":    footer.TwSrec,
	})
}
