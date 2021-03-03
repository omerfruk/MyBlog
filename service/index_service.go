package service

import (
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateTopBar(logo string, option [4]string) {
	temp := models.Topbar{
		Logo:   logo,
		Option: option,
	}
	if tempCont:= database.DB().Where("logo = ? " , temp.Logo).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}

func CreateEntry(img string, header string, text string, buton string) {
	temp := models.Entry{
		ImgSrc:     img,
		Header:     header,
		Text:       text,
		ButtonText: buton,
	}
	if tempCont:= database.DB().Where("header = ? and text = ?" ,temp.Header,temp.Text ).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}
func CreateInstructions(title string, info string, opt [3]string, desc [3]string) {
	temp:=models.Instructions{
		Title:        title,
		Info:         info,
		Option:       opt,
		Descriptions: desc,
	}
	if tempCont:= database.DB().Where("title = ? and info = ?" ,temp.Title,temp.Info).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}
func CreatePortfolio(imgsrc string,title string, desc string)  {
	temp:=models.Portfolio{
		ImgSrc:       imgsrc,
		Title:        title,
		Descriptions: desc,
	}
	if tempCont:= database.DB().Where("imgsrc = ? " , temp.ImgSrc).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}
