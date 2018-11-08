package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Consumer struct {
	Id               string    `orm:"pk" json:"id"`
	Name             string    `orm:"column(name)" json:"name"   form:"name"`
	PhoneNumber       string    `orm:"column(phone_number)" json:"phoneNumber"    form:"phoneNumber"`
	ConsumerType      string    `orm:"column(consumer_type)" json:"consumerType" form:"consumerType"`
	Deposit   string            `orm:"column(deposit)" json:"deposit"  form:"deposit"`
	CreatedAt        time.Time  `orm:"column(created_at)" json:"createdAt" `
	ElectricityFee    float64   `orm:"column(electricity_fee)" json:"electricityFee"   form:"electricityFee"`
	CustodianFee      float64   `orm:"column(custodian_fee)" json:"custodianFee"   form:"custodianFee"`
	Source            float64   `orm:"column(source)" json:"source"   form:"source"`
	PayWay            float64   `orm:"column(pay_way)" json:"payWay"   form:"payWay"`
}


func init()  {
	orm.RegisterModel(new(Consumer))
}

func (consumer *Consumer) GetById() *Consumer {
	o := orm.NewOrm()
	o.Read(consumer,"id")
	return consumer
}

func GetConsumers(typeId string, bdcId string)  {
	o := orm.NewOrm()
	var consumer []*Consumer
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("t1.name,t1.id").From("consumer_goods t ").
		LeftJoin("consumer t1").On("t.consumer_id = t1.id").
		Where(" t.bdc_id = ?").And(" t.type_id = ?").And("t.wait_join_count !=0").GroupBy("t1.id")
	sql := qb.String()
	o.Raw(sql,typeId,bdcId).QueryRows(&consumer)
}
