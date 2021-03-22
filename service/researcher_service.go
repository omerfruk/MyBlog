package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func GetResearch(area string) []models.Research {
	var research []models.Research
	err := database.DB().Where("area = ?", area).Find(&research).Error
	if err != nil {
		fmt.Println(err)
	}
	return research
}

func CreateResearch(research models.Research) {
	err := database.DB().Where("title=? and text=?", research.Title, research.Text).First(&research).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&research).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bir arastirma bulunuyor")
	}

}

func DeleteResearch(title string) {
	temp := new(models.Research)
	database.DB().Where("title=?", title).Delete(&temp)
}

func UpdateResarch(id int, area string, title string, text string) {
	temp := new(models.Research)
	database.DB().Where("id=?", id).Find(&temp)
	temp.Area = area
	temp.Text = text
	temp.Title = title
	database.DB().Save(&temp)
}
