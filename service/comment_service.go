package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func CommentCreate(name string, phone string, mail string, title string, text string, img string) {
	temp := models.Comment{
		Fullname: name,
		Phone:    phone,
		Mail:     mail,
		Title:    title,
		Text:     text,
		Imgsrc:   img,
	}
	err := database.DB().Where("fullname = ? and title = ?", temp.Fullname, temp.Title).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&temp).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boytle bi kayıt var")
	}

}
func GetAllComments() []models.Comment {
	var temp []models.Comment
	database.DB().Find(&temp)

	return temp
}
