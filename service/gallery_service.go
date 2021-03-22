package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

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

func GetFotoAll() []models.FotoGallery {
	var temp []models.FotoGallery
	database.DB().Find(&temp)
	return temp
}
