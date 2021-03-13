package models

import (
	"gorm.io/gorm"
)

type Topbar struct {
	gorm.Model
	Logo    string
	Home    string
	Future  string
	Port    string
	Contact string
	Login   string
}

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestSingUp struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
	Title     string `json:"title"`
	Info      string `json:"info"`
	LeftIntro string `json:"left_intro"`
	MidIntro  string `json:"mid_intro"`
	RghtIntro string `json:"rght_intro"`
	LeftDesc  string `json:"left_desc"`
	MidDesc   string `json:"mid_desc"`
	RghtDesc  string `json:"rght_desc"`
}

type AdminPage struct {
	Topbar Topbar
	Info   User
	User   []User
}

type FotoStruct struct {
	FotoGallery []FotoGallery
	FooterBar   FooterBar
}
type Anasayfa struct {
	Portfolio []Portfolio
	Entry     Entry
	Topbar    Topbar
	Intro     Instructions
	Footer    FooterBar
}

type Info struct {
	Topbar Topbar
	User   User
	Footer FooterBar
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
type FotoGallery struct {
	gorm.Model
	ImgSrc string `json:"img_src"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}
type Research struct {
	gorm.Model
	Area  string `json:"area"`
	Title string `json:"title"`
	Text  string `json:"text"`
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
	ImgSrc      string    `json:"img_src"`
	Mail        string    `json:"mail"`
	Password    string    `json:"password"`
	Header      string    `json:"header"`
	Fullname    string    `json:"fullname"`
	Information string    `json:"information"`
	Authority   Authority `json:"authority"`
}
