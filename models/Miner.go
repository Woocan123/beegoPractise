package models

import "github.com/astaxie/beego/orm"

type Miner struct {
	Id           string `orm:"pk" json:"id"`
	MineWorkerNo string `orm:"column(mine_worker_no)" json:"mineWorkerNo"   form:"mineWorkerNo"`
	ConsumerId   string `orm:"column(consumer_id)" json:"consumerId"    form:"consumerId"`
	MinePool     string `orm:"column(mine_pool)" json:"minePool" form:"minePool"`
}

func init()  {
	orm.RegisterModel(new(Miner))
}

func (miner *Miner) GetById() *Miner  {
	o :=orm.NewOrm()
	o.Read(miner,"id")
	return miner
}