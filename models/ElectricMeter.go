package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ElectricMeterReading struct {
	Id string                   `orm:"pk" json:"id"`
	ElectricMeterReading string `orm:"column(electric_meter_reading)" json:"electricMeterReading" form:"electricMeterReading"`
	Type string                 `orm:"column(type)" json:"type" form:"type"`
	EngineRoomNumber string     `orm:"column(engine_room_number)" json:"engineRoomNumber" form:"engineRoomNumber"`
	RackNumber string           `orm:"column(rack_number)" json:"rackNumber" form:"engineRoomNumber"`
	RecordTime time.Time        `orm:"column(record_time)" json:"recordTime" form:"recordTime"`
	CreatedAt time.Time         `orm:"column(created_at)" json:"createdAt" form:"createdAt"`
	CreateUser string           `orm:"column(create_user)" json:"createUser" form:"createUser"`
	BdcId string                `orm:"column(bdc_id)" json:"bdcId" form:"bdcId"`
	Status string               `orm:"column(status)" json:"status" form:"status"`
	ElectricMeter *ElectricMeter `orm:"rel(one)"`
	//ElectricMeterId string      `orm:"column(electric_meter_id)" json:"electricMeterId" form:"electricMeterId"`
	//Photo *Photo                `orm:"rel(one)"`
}

type ElectricMeter struct {
	Id string `orm:"pk" json:"id"`
	BdcId string `orm:"column(bdc_id)" json:"bdcId" form:"bdcId"`
	Name string `orm:"column(name)" json:"name" form:"name"`
	ElectricMeterReading *ElectricMeterReading `orm:"reverse(one)"`
}


func init()  {
	orm.RegisterModel(new(ElectricMeterReading),new(ElectricMeter))
}

func (electricMeter *ElectricMeterReading) AddOrUpdate() int64 {
	o := orm.NewOrm()
	num,_ := o.InsertOrUpdate(electricMeter,"id")
	return num
}

func (electricMeter *ElectricMeterReading) GetById() *ElectricMeterReading {
	o := orm.NewOrm()
	o.Read(electricMeter,"id")
	if electricMeter.ElectricMeter != nil {
		o.Read(electricMeter.ElectricMeter)
	}
	return electricMeter
}