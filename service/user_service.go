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
	err := database.DB().Where("authority=?", 2).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}
func GetUserAll() models.User {
	var user models.User
	err := database.DB().First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}
func SingUPUser(phone string, name string, email string, pass string) {
	temp := new(models.User)
	temp.Information = phone
	temp.Fullname = name
	temp.Mail = email
	temp.Password = pass
	temp.Authority = models.Basic
	if UserTemp := database.DB().Where("email=?", temp.Mail).First(&temp); UserTemp.Error != nil {
		database.DB().Create(&temp)
	}
}
func CreateUser(img string, header string, name string, mail string, info string, isAdmin bool) {
	temp := new(models.User)
	temp.Header = header
	temp.ImgSrc = img
	temp.Mail = mail
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
func DeleteUser(id string) {
	temp := new(models.User)
	database.DB().Where("id=?", id).Delete(&temp)
}
func UpdateUser(id uint, mail string, name string, info string, img string) {

	var temp models.User

	database.DB().Where("id=?", id).Find(&temp)
	temp.Information = info
	temp.Fullname = name
	temp.ImgSrc = img
	temp.Mail = mail
	database.DB().Save(&temp)

}
