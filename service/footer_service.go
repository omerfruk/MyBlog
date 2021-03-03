package service

import (
	"github.com/omerfruk/my-blog/database"
	"github.com/omerfruk/my-blog/models"
)

func CreateFooter(imgsrc string,name string,emp string,inst string,fac string,git string,tw string)  {
	temp:=models.FooterBar{
		ImgSrc:     imgsrc,
		FullName:   name,
		Employment: emp,
		InstaSrc:   inst,
		FaccSrc:    fac,
		GitSrc:     git,
		TwSrec:     tw,
	}
	if tempCont:= database.DB().Where("img_src = ? " , temp.ImgSrc).First(&temp);tempCont.Error != nil{
		database.DB().Create(&temp)
	}
}
