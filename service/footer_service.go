package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func CreateFooter(bar models.FooterBar) {
	err := database.DB().Where("img_src = ? ", bar.ImgSrc).First(&bar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&bar).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bir top bar vardir")
	}
}
func GetFooter(name string) models.FooterBar {
	var temp models.FooterBar
	err := database.DB().Where("full_name = ? ", name).First(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}
func DeleteFooter(name string) {
	var temp models.FooterBar
	database.DB().Where("full_name = ? ", name).Delete(&temp)
}
func UpdateFooter(oldName string, imgsrc string, name string, emp string, inst string, fac string, git string, tw string) {
	var temp models.FooterBar
	database.DB().Where("full_name = ?", oldName).Find(&temp)
	temp.ImgSrc = imgsrc
	temp.FullName = name
	temp.Employment = emp
	temp.InstaSrc = inst
	temp.FaccSrc = fac
	temp.GitSrc = git
	temp.TwSrec = tw
	database.DB().Save(&temp)
}
