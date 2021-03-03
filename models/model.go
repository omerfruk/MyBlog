package models

import (
	"gorm.io/gorm"
)

type Topbar struct {
	gorm.Model
	Logo   string
	Home string
	Future string
	Port string
	Contact string
}

type Entry struct {
	gorm.Model
	ImgSrc     string `json:"img_src"`
	Header     string `json:"header"`
	Text       string `json:"text"`
	ButtonText string `json:"button_text"`
}
type Instructions struct {
	gorm.Model
	Title string `json:"title"`
	Info  string `json:"info"`
	Intro []intro
}
type intro struct {
	gorm.Model
	optinal        string
	desc           string
	InstructionsId uint8
}
type Portfolio struct {
	gorm.Model
	ImgSrc       string `json:"img_src"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
}

type FooterBar struct {
	gorm.Model
	ImgSrc     string `json:"img_src"`
	FullName   string `json:"full_name"`
	Employment string `json:"employment"`
	InstaSrc   string `json:"insta_src"`
	FaccSrc    string `json:"facc_src"`
	GitSrc     string `json:"git_src"`
	TwSrec     string `json:"tw_srec"`
}
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
