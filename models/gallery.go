package models

import (
	"gorm.io/gorm"
)

type FotoGallery struct {
	gorm.Model
	ImgSrc string `json:"img_src"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type FotoStruct struct {
	FotoGallery []FotoGallery
	FooterBar   FooterBar
}
