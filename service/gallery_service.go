package service

import (
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateFoto(imgsrc string, title string, text string) {
	temp := models.FotoGallery{
		ImgSrc: imgsrc,
		Title:  title,
		Text:   text,
	}
	if tempFoto := database.DB().Where("img_src = ? ", imgsrc).First(&temp); tempFoto.Error != nil {
		database.DB().Create(&temp)
	}

}
