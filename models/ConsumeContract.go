package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ConsumerContract struct {
	Id        string                  `orm:"pk" json:"id"`
	Name      string                  `orm:"column(name)" json:"name"   form:"name"`
	ConsumerId      string            `orm:"column(consumer_id)" json:"consumerId"    form:"consumerId"`
	PaymentType   string              `orm:"column(payment_type)" json:"paymenttype" form:"paymenttype"`
	ElectricityFee string             `orm:"column(electricity_fee)" json:"electricityFee"  form:"electricityFee"`
	CustodianFee  string              `orm:"column(custodian_fee)" json:"custodianFee"  form:"custodianFee"`
	DepositStartTime  string          `orm:"column(deposit_start_time)" json:"depositStartTime"  form:"depositStartTime"`
	DepositEndTime   string           `orm:"column(deposit_end_time)" json:"depositEndTime"  form:"depositEndTime"`
	Status string  					  `orm:"column(status)" json:"status"`
	CreatedAt  time.Time 			  `orm:"column(created_at)" json:"createdAt" `
	PutawayFee float64 				  `orm:"column(putaway_fee)" json:"putawayFee"   form:"putawayFee"`
	DepositServeFee float64 		  `orm:"column(deposit_serve_fee)" json:"depositServeFee"   form:"depositServeFee"`
	ElecForegift float64 			  `orm:"column(elec_foregift)" json:"elecForegift"   form:"elecForegift"`
	SeatBail float64 				  `orm:"column(seat_bail)" json:"seatBail"   form:"seatBail"`
	TotalAmount float64 			  `orm:"column(total_amount)" json:"totalAmount"   form:"totalAmount"`
}

func init()  {
	orm.RegisterModel(new(ConsumerContract))
}

func (consumerContract *ConsumerContract) GetById() *ConsumerContract {
	o := orm.NewOrm()
	o.Read(consumerContract,"id")
	return consumerContract
}
