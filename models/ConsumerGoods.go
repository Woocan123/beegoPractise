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


func init()  {
	orm.RegisterModel(new(ConsumerGoods))
}

func GetMachineGoods(typeId string, bdcId string, consumerId string, minerNo string) []*ConsumerGoods {
	o := orm.NewOrm()
	var consumerGoods []*ConsumerGoods
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("t3.miner_no as minerNo, if((t3.wait_join_count - t3.defective_count) < 0,t3.wait_join_count," +
		"t3.wait_join_count - t3.defective_count) as waitPutawayCount,t3.runCount, t3.putawayCount, t3.joinrepertoryCount, " +
		"t3.check_count as checkCount, t3.wait_join_count as waitJoinCount, t3.defective_count as defectiveCount, t3.undercarriageCount, " +
		"t3.id as consumerGoodsId, t4.id,t4.name,if( t6.pass_count is null, 0,t6.pass_count) as passCount,if( t6.bad_count is null, 0,t6.bad_count) as badCount, " +
		"t7.counts as repairCount, t8.counts as deliveryCount,t4.spec as spec").
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