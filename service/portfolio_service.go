package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func CreatePortfolio(portfolio models.Portfolio) {
	err := database.DB().Where("img_src = ? ", portfolio.ImgSrc).First(&portfolio).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&portfolio).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bir portfolio bulunmakta")
	}
}

func GetPortfolio() []models.Portfolio {
	var temp []models.Portfolio
	err := database.DB().Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}

func DeletePortfolio(title string) {
	var temp models.Portfolio
	database.DB().Where("title = ?", title).Delete(&temp)
}
func UpdatePortfolio(oldTitle string, imgsrc string, title string, desc string) {
	var temp models.Portfolio
	database.DB().Where("title = ?", oldTitle).Find(&temp)
	temp.ImgSrc = imgsrc
	temp.Title = title
	temp.Descriptions = desc
	database.DB().Save(&temp)
}
