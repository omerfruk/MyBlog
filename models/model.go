package models

import (
	"gorm.io/gorm"
)

type Research struct {
	gorm.Model
	Area  string `json:"Area"`
	Title string `json:"Title"`
	Text  string `json:"Text"`
}

type ReserchStr struct {
	ResearchStr []Research `json:"research"`
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

