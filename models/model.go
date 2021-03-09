package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
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

type Config struct {
	// Allowed session duration
	// Optional. Default value 24 * time.Hour
	Expiration time.Duration

	// Storage interface to store the session data
	// Optional. Default value memory.New()
	Storage fiber.Storage

	// Name of the session cookie. This cookie will store session key.
	// Optional. Default value "session_id".
	CookieName string

	// Domain of the CSRF cookie.
	// Optional. Default value "".
	CookieDomain string

	// Path of the CSRF cookie.
	// Optional. Default value "".
	CookiePath string

	// Indicates if CSRF cookie is secure.
	// Optional. Default value false.
	CookieSecure bool

	// Indicates if CSRF cookie is HTTP only.
	// Optional. Default value false.
	CookieHTTPOnly bool

	// Indicates if CSRF cookie is HTTP only.
	// Optional. Default value false.
	CookieSameSite string

	// KeyGenerator generates the session key.
	// Optional. Default value utils.UUID
	KeyGenerator func() string
}
