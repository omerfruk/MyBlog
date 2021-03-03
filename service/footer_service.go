package service

import (
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateFooter(imgsrc string, name string, emp string, inst string, fac string, git string, tw string) {
	temp := models.FooterBar{
		ImgSrc:     imgsrc,
		FullName:   name,
		Employment: emp,
		InstaSrc:   inst,
		FaccSrc:    fac,
		GitSrc:     git,
		TwSrec:     tw,
	}
	if tempCont := database.DB().Where("img_src = ? ", temp.ImgSrc).First(&temp); tempCont.Error != nil {
		database.DB().Create(&temp)
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
