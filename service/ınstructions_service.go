package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateInstructions(instructions models.Instructions) {
	err := database.DB().Where("title = ? and info = ?", instructions.Title, instructions.Info).First(&instructions).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&instructions).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle bir giris bulunmaktadir")
	}
}
func DeleteInstructions(title string) {
	var temp models.Instructions
	database.DB().Where("title = ? ", title).Delete(&temp)
}
func GetInstructions(title string) models.Instructions {
	var temp models.Instructions
	err := database.DB().Where("title = ? ", title).Find(&temp).Error
	if err != nil {
		fmt.Println(err)
	}
	return temp
}
func UpdateInstructions(oldTitle string, title string, info string, LeftOpt string, MidOpt string, RghtOpt string, LeftDesc string, MidDesc string, RghtDesc string) {
	var temp models.Instructions
	database.DB().Where("title = ?", oldTitle).Find(&temp)

	temp.Title = title
	temp.Info = info
	temp.LeftIntro = LeftOpt
	temp.MidIntro = MidOpt
	temp.RghtIntro = RghtOpt
	temp.LeftDesc = LeftDesc
	temp.RghtDesc = RghtDesc
	temp.MidDesc = MidDesc

	database.DB().Save(&temp)
}
