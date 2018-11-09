package models

import "github.com/astaxie/beego/orm"

type Miner struct {
	Id           string `orm:"pk" json:"id"`
	MineWorkerNo string `orm:"column(mine_worker_no)" json:"mineWorkerNo"   form:"mineWorkerNo"`
	ConsumerId   string `orm:"column(consumer_id)" json:"consumerId"    form:"consumerId"`
	MinePool     string `orm:"column(mine_pool)" json:"minePool" form:"minePool"`
}

//里面的字段不能与上面Miner内的字段一样，否则查询不出
type MinerO struct {
	Mine string
	ConsumerGoods []*ConsumerGoods `orm:="-"`
}


func init()  {
	orm.RegisterModel(new(Miner))
}

func (miner *Miner) GetById() *Miner  {
	o :=orm.NewOrm()
	o.Read(miner,"id")
	return miner
}

func GetMiners(typeId string,bdcId string,consumerId string) []*MinerO  {
	var miner []*MinerO
	o := orm.NewOrm()
	qb,_ :=orm.NewQueryBuilder("mysql")
	qb.Select("t.miner_no mine").From("consumer_goods t").
		Where(" t.bdc_id = ?").And(" t.type_id = ?").And("t.consumer_id = ?").GroupBy("t.miner_no")
	sql := qb.String()
	o.Raw(sql,bdcId,typeId,consumerId).QueryRows(&miner)
	return miner
}