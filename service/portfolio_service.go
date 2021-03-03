package service

import (
	"fmt"
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

func GetPortfolio() []models.Portfolio {
	var temp []models.Portfolio
	err := database.DB().Find(&temp).Error
	if err != nil{
		fmt.Println(err)
	}
	return temp
}
func DeletePortfolio(title string)  {
	var temp models.Portfolio
	database.DB().Where("title = ?",title).Delete(&temp)
}
func UpdatePortfolio(oldTitle string,imgsrc string,title string, desc string)  {
	var temp models.Portfolio
	database.DB().Where("title = ?" , oldTitle).Find(&temp)
	temp.ImgSrc=imgsrc
	temp.Title=title
	temp.Descriptions=desc
	database.DB().Save(&temp)
}