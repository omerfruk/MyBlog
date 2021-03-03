package service

import (
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func GetUser(name string) models.User {
	var user models.User
	err := database.DB().Where("fullname=?", name).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func CreateUser(img string,header string, name string, info string, isAdmin bool) {
	temp := new(models.User)
	temp.Header = header
	temp.ImgSrc=img
	temp.Fullname = name
	temp.Information = info
	if isAdmin == true {
		temp.Authority = models.Admin
	} else {
		temp.Authority = models.Basic
	}
	if UserTemp := database.DB().Where("fullname=?", temp.Fullname).First(&temp); UserTemp.Error != nil {
		database.DB().Create(&temp)
	}
}
func DeleteUser(name string) {
	temp := new(models.User)
	database.DB().Where("fullname=?", name).Delete(&temp)
}
func UpdateUser(oldName string,header string, name string, info string)  {
	var temp models.User
	database.DB().Where("fullname=?",oldName).Find(&temp)
	temp.Information=info
	temp.Fullname=name
	temp.Header=header
	database.DB().Save(&temp)

}
