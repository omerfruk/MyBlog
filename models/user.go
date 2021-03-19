package models

import (
	"gorm.io/gorm"
)

type Authority int

const (
	Basic Authority = iota + 1
	Admin
)

type User struct {
	gorm.Model
	ImgSrc      string    `json:"img_src"`
	Mail        string    `json:"mail"`
	Password    string    `json:"password"`
	Header      string    `json:"header"`
	Fullname    string    `json:"fullname"`
	Information string    `json:"information"`
	Authority   Authority `json:"authority"`
}
