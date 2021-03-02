package models

import (
	"gorm.io/gorm"
)

type Research struct {
	gorm.Model
	Area  string `json:"area"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Authority int

const (
	Basic Authority = iota + 1
	Admin
)

type User struct {
	gorm.Model
	Header      string    `json:"header"`
	Fullname    string    `json:"fullname"`
	Information string    `json:"information"`
	Authority   Authority `json:"authority"`
}
