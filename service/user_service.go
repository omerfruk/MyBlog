package service

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

type userService struct {
	db *gorm.DB
}

func GetUser(name string) models.User {
	var user models.User
	err:=database.DB().Where("fullname=?",name).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func CreateUser(header string,name string, info string, isAdmin bool) {
	temp := new(models.User)
	temp.Header=header
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
