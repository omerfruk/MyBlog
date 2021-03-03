package service

import (
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreatePortfolio(imgsrc string,title string, desc string)  {
	temp:=models.Portfolio{
		ImgSrc:       imgsrc,
		Title:        title,
		Descriptions: desc,
	}
	if tempCont:= database.DB().Where("img_src = ? " , temp.ImgSrc).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}