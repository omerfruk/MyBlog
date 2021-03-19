package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Imgsrc   string `json:"imgsrc"`
}
