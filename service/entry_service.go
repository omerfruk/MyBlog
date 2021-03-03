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
func GetEntry(header string)models.Entry{
	var entry models.Entry
	err:=database.DB().Where("header = ?",header).First(&entry)
	if err !=nil{
		fmt.Println(err)
	}
	return entry
}
func DeleteEntry(header string)  {
	temp:=new(models.Entry)
	database.DB().Where("header = ?", header).Delete(&temp)
}
func UpdateEntry(oldHeader string,img string,header string,text string,button string)  {
	var temp models.Entry
	database.DB().Where("header = ?",header).Find(&temp)
	temp.ImgSrc=img
	temp.Header=header
	temp.Text=text
	temp.ButtonText=button
	database.DB().Save(&temp)
}
