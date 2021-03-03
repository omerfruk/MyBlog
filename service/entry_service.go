package service

import (
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateEntry(img string, header string, text string, buton string) {
	temp := models.Entry{
		ImgSrc:     img,
		Header:     header,
		Text:       text,
		ButtonText: buton,
	}
	if tempCont:= database.DB().Where("header = ? and text = ?" ,temp.Header,temp.Text ).First(&temp);tempCont.Error != nil{
		err:=database.DB().Create(&temp)
		if err !=nil{
			fmt.Println(err)
		}
	}
}
