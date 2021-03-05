package service

import (
	"fmt"
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)
func GetResearch(area string) [] models.Research{
	var research [] models.Research
	err := database.DB().Where("area =?", area).Find(&research).Error
	if err != nil{
		fmt.Println(err)
	}
	return research
}

func CreateResearch(area string, title string, text string) {
	temp := new(models.Research)
	temp.Area = area
	temp.Text = text
	temp.Title = title

	if researchTemp := database.DB().Where("title=? and text=?", temp.Title, temp.Text).First(&temp); researchTemp.Error != nil {
		database.DB().Create(&temp)
	}
}

func DeleteResearch(title string)  {
	temp:=new(models.Research)
	database.DB().Where("title=?",title).Delete(&temp)
}

func UpdateResarch(id int,area string,title string, text string)  {
	temp:=new(models.Research)
	database.DB().Where("id=?",id).Find(&temp)
	temp.Area=area
	temp.Text=text
	temp.Title=title
	database.DB().Save(&temp)
}