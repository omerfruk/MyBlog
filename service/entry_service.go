package service

import (
	"errors"
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
	"gorm.io/gorm"
)

func CreateEntry(entry models.Entry) {
	err := database.DB().Where("header = ? and text = ?", entry.Header, entry.Text).First(&entry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&entry).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bilgi database de bulunmakta")
	}
}
func GetEntry(header string) models.Entry {
	var entry models.Entry
	err := database.DB().Where("header = ?", header).First(&entry).Error
	if err != nil {
		fmt.Println(err)
	}
	return entry
}
func DeleteEntry(header string) {
	temp := new(models.Entry)
	database.DB().Where("header = ?", header).Delete(&temp)
}
func UpdateEntry(oldHeader string, img string, header string, text string, button string) {
	var temp models.Entry
	database.DB().Where("header = ?", header).Find(&temp)
	temp.ImgSrc = img
	temp.Header = header
	temp.Text = text
	temp.ButtonText = button
	database.DB().Save(&temp)
}
