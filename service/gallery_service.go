package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func GetFotoAll() []models.FotoGallery {
	var temp []models.FotoGallery
	database.DB().Find(&temp)
	return temp
}

func CreateFoto(gallery models.FotoGallery) {
	err := database.DB().Where("img_src = ? ", gallery.ImgSrc).First(&gallery).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = database.DB().Create(&gallery).Error
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("boyle uzantıda bir resim bulunmaktadir")
	}
}

func UpdateFoto(img string, gallery models.FotoGallery) models.FotoGallery {
	var temp models.FotoGallery
	err := database.DB().Where("img_src = ?", img).Find(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("boyle bir resim yok")
		return gallery
	}
	temp = gallery
	return temp
}

func DeleteFoto(img string) {
	var temp models.FotoGallery
	database.DB().Where("img_src = ?", img).Delete(&temp)
}
