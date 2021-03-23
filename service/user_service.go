package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func GetUserMail(mail interface{}) models.User {
	var user models.User
	err := database.DB().Where("mail = ?", mail).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}
func GetUserId(id string) models.User {
	var user models.User
	err := database.DB().Where("id = ?", id).First(&user).Error
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

func GetAdmin(name string) models.User {
	var temp models.User
	err := database.DB().Where("fullname = ? and authority = ?", name, models.Admin).Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}

func CreateUser(user models.User, isAdmin bool) {
	if isAdmin == true {
		user.Authority = models.Admin
	} else {
		user.Authority = models.Basic
	}
	user.Password = Sha256String(user.Password)
	err := database.DB().Where("mail=?", user.Mail).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&user).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("bu maile ait bir hesab zaten var")
	}
}

func UpdateUser(user models.User, Olduser models.User) {
	Olduser = user
	database.DB().Save(&Olduser)
}

func DeleteUser(id string) {
	var temp models.User
	database.DB().Where("id=?", id).Delete(&temp)
}

func Find() ([]models.User, error) {
	var user []models.User
	err := database.DB().Find(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
	}
	return user, err
}
