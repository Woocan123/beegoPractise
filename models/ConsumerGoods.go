package models

import (
	"github.com/astaxie/beego/orm"
)

type ConsumerGoods struct {
	Id                  string    `orm:"pk" json:"id"`
	GoodsId        		string    `orm:"column(goods_id)" json:"goodsId"`
	TypeId         		string    `orm:"column(type_id)" json:"typeId" `
	SpecId         		string    `orm:"column(spec_id)" json:"specId" `
	ConsumerId     		string    `orm:"column(consumer_id)" json:"consumerId"  `
	ContractId     		string    `orm:"column(contract_id)" json:"contractId"`
	MinerNo        		string    `orm:"column(miner_no)" json:"minerNo"`
	BdcId          		string    `orm:"column(bdc_id)" json:"bdcId"`
	CheckCount          int64   `orm:"column(source)" json:"source" `
	RunCount            int64   `orm:"column(pay_way)" json:"payWay" `
	Remark              string   `orm:"column(remark)" json:"remark" `
	PutawayCount        int64   `orm:"column(putawayCount)" json:"putawayCount" `
	UndercarriageCount  int64   `orm:"column(undercarriageCount)" json:"undercarriageCount" `
	JoinrepertoryCount  int64   `orm:"column(joinrepertoryCount)" json:"joinrepertoryCount" `
	WaitJoinCount       int64   `orm:"column(wait_join_count)" json:"waitJoinCount" `
	DefectiveCount      int64   `orm:"column(defective_count)" json:"defectiveCount" `
}

//用构造查询的时候struct必须要有orm，不然查询出来的字段无法与struct里的字段对应
type ConsumerGoodsO struct {
	Id                 string  `orm:"pk" json:"id"`
	MinerNo            string	`orm:"column(minerNo)"`
	WaitPutawayCount   int64	`orm:"column(waitPutawayCount)"`
	RunCount           int64
	PutawayCount       int64	`orm:"column(putawayCount)"`
	JoinrepertoryCount int64	`orm:"column(joinrepertoryCount)"`
	CheckCount         int64	`orm:"column(checkCount)"`
	WaitJoinCount      int64	`orm:"column(waitJoinCount)"`
	DefectiveCount     int64	`orm:"column(defectiveCount)"`
	UndercarriageCount int64	`orm:"column(undercarriageCount)"`
	ConsumerGoodsId    string	`orm:"column(consumerGoodsId)"`
	Name               string
	PassCount          int64	`orm:"column(passCount)"`
	BadCount           int64	`orm:"column(badCount)"`
	RepairCount        int64	`orm:"column(repairCount)"`
	DeliveryCount      int64	`orm:"column(deliveryCount)"`
	Spec               string
}



func init()  {
	orm.RegisterModel(new(ConsumerGoods),new(ConsumerGoodsO))
}

func GetMachineGoods(typeId string, bdcId string, consumerId string, minerNo string) []*ConsumerGoodsO {
	o := orm.NewOrm()
	var consumerGoods []*ConsumerGoodsO
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("t3.miner_no minerNo, if((t3.wait_join_count - t3.defective_count) < 0,t3.wait_join_count," +
		"t3.wait_join_count - t3.defective_count) waitPutawayCount,t3.runCount, t3.putawayCount, t3.joinrepertoryCount, " +
		"t3.check_count checkCount, t3.wait_join_count waitJoinCount, t3.defective_count defectiveCount, t3.undercarriageCount, " +
		"t3.id consumerGoodsId, t4.id,t4.name,if( t6.pass_count is null, 0,t6.pass_count) passCount,if( t6.bad_count is null, 0,t6.bad_count) badCount, " +
		"t7.counts repairCount, t8.counts deliveryCount,t4.spec spec").
		From("consumer_goods t3").
		LeftJoin("goods_name t4").On("t3.goods_id=t4.id").
		LeftJoin("repertory t6").On("t3.id = t6.consumer_goods_id").
		LeftJoin("machine_repair_detail t7").On("t3.id = t7.consumer_goods_id").
		LeftJoin("machine_delivery_detail t8").On("t3.id = t8.consumer_goods_id").
		Where(" t3.bdc_id = ?").And(" t3.type_id = ?").And("t3.consumer_id = ?").And("t3.miner_no = ?").GroupBy("t3.id")
	sql := qb.String()
	o.Raw(sql,bdcId,typeId,consumerId,minerNo).QueryRows(&consumerGoods)
	return consumerGoods
}