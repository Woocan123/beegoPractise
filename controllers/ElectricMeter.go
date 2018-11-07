package controllers

import
(
	"fmt"
	"github.com/astaxie/beego"
. "beegoPractise/models"
	"github.com/gofrs/uuid"
	"strings"
)

type ElectricMeterController struct {
	beego.Controller
}

func (e *ElectricMeterController) GetById()  {
	id :=e.GetString("id")
	var electric = new(ElectricMeterReading)
	electric.Id = id
	result :=electric.GetById()
	e.Data["json"] = map[string]interface{}{"msg":"成功","data":result,"code":0}
	e.ServeJSON()
}

func (e *ElectricMeterController) AddOrUpdate()  {
	id :=uuid.Must(uuid.NewV4()).String()
	id = strings.Replace(id,"-","",-1)
	file,fileheader,_ := e.GetFile("photos")
	a := file
	b := fileheader
	fmt.Print(a,b)
	var electric ElectricMeterReading
	if err := e.ParseForm(&electric); err != nil{

	}
	electric.Id = id
	electric.AddOrUpdate()
	e.Data["json"] = map[string]interface{}{"msg":"成功","code":0}
	e.ServeJSON()
}



