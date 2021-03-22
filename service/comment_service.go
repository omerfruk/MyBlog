package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func CommentCreate(comment models.Comment) {
	err := database.DB().Where("fullname = ? and title = ?", comment.Fullname, comment.Title).First(&comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&comment).Error
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
