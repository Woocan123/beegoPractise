package models

import "github.com/astaxie/beego/orm"

type Photo struct {
	Id string `orm:"pk" json:"id"`
	BusinessId string `orm:"column(business_id)" json:"businessId"`
	Path string `orm:"column(path)" json:"path"`
	FileName string `orm:"column(file_name)" json:"fileName"`
	//ElectricMeterReading *ElectricMeterReading `orm:"reverse(one)"`
}

func init()  {
	orm.RegisterModel(new(Photo))
}
