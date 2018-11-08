package controllers

import (
	. "beegoPractise/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

type ElectricMeterController struct {
	beego.Controller
}

func (e *ElectricMeterController) GetById()  {
	id :=e.GetString("id")
	var electric = new(ElectricMeterReadingRel)
	electric.Id = id
	result :=electric.GetById()
	e.Data["json"] = map[string]interface{}{"msg":"成功","data":result,"code":0}
	e.ServeJSON()
}

func (e *ElectricMeterController) AddOrUpdate()  {
	id :=uuid.Must(uuid.NewV4()).String()
	id = strings.Replace(id,"-","",-1)
	f,h,_ := e.GetFile("photos")
	defer f.Close()
	path := fmt.Sprintf("%s/trust/uploads/electricMeterReading/%s",beego.AppConfig.String("root"),fmt.Sprintf("%s.jpeg",uuid.Must(uuid.NewV4()).String()))
	e.SaveToFile("photos",fmt.Sprintf("%s/%s",beego.AppConfig.String("path"),path))
	var electric ElectricMeterReading
	if err := e.ParseForm(&electric); err != nil{

	}
	electric.Id = id
	electric.CreatedAt = time.Now()
	recordTime,_ := time.Parse("2006/01/02 15:04:05","2018/11/07 12:00:00")
	electric.RecordTime = time.Unix(recordTime.Unix(),0).Local()
	electric.AddOrUpdate()
	var photo = new(Photo)
	photoId :=uuid.Must(uuid.NewV4()).String()
	photoId = strings.Replace(photoId,"-","",-1)
	photo.Id = photoId
	photo.BusinessId = id
	photo.Path = path
	photo.FileName = h.Filename
	photo.Add()
	e.Data["json"] = map[string]interface{}{"msg":"成功","code":0}
	e.ServeJSON()
}



