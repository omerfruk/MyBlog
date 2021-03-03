package service

import (
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateTopBar(logo string, home string, future string, port string, contact string) {
	temp := models.Topbar{
		Logo:    logo,
		Home:    home,
		Future:  future,
		Port:    port,
		Contact: contact,
	}
	if tempCont := database.DB().Where("Logo = ? ", temp.Logo).First(&temp); tempCont.Error != nil {
		database.DB().Create(&temp)
	}
}
func DeleteTopBar(logo string) {
	var temp models.Topbar
	database.DB().Where("logo = ?", logo).Delete(&temp)
}

func GetTopBar(logo string) models.Topbar {
	var temp models.Topbar
	err := database.DB().Where("logo = ? ", logo).Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}

func UpdateTopBar(oldLogo string,logo string, home string, future string, port string, contact string)  {
	var temp models.Topbar
	database.DB().Where("logo = ?",oldLogo).Find(&temp)

	temp.Logo=logo
	temp.Home=home
	temp.Future=future
	temp.Port=port
	temp.Contact=contact

	database.DB().Save(&temp)
}
