package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func GetTopBar(logo string) models.Topbar {
	var temp models.Topbar
	err := database.DB().Where("logo = ? ", logo).Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp

}

func CreateTopBar(topbar models.Topbar) {
	err := database.DB().Where("Logo = ? ", topbar.Logo).First(&topbar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&topbar).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bir tobbar mevcud")
	}

}

func UpdateTopBar(oldLogo string, logo string, home string, future string, port string, contact string, log string) {
	var temp models.Topbar
	database.DB().Where("logo = ?", oldLogo).Find(&temp)
	temp.Logo = logo
	temp.Home = home
	temp.Future = future
	temp.Port = port
	temp.Contact = contact
	temp.Login = logo

	database.DB().Save(&temp)
}

func DeleteTopBar(logo string) {
	var temp models.Topbar
	database.DB().Where("logo = ?", logo).Delete(&temp)
}
