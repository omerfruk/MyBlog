package structure

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Title string
	User string
	Information string
}