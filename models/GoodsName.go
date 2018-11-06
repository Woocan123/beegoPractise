package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type GoodsName struct {
	Id        string `orm:"pk"`
	TypeId    string `orm:"column(type_id)"`
	Name      string `orm:"column(name)"`
	Spec      string `orm:"column(spec)"`
	Version   string `orm:"column(version)"`
	CreatedAt time.Time `orm:"column(created_at)"`
	Hashrate  string `orm:"column(hashrate)"`
	Watt      string `orm:"column(watt)"`
	TopWatt   string `orm:"column(top_watt)"`
	WuhaiWatt string `orm:"column(wuhai_watt)"`
	CoinType  string `orm:"column(coin_type)"`
}
//type GoodsName struct {
//	Id        string `orm:"pk"`
//	TypeId    string
//	Name      string
//	Spec      string
//	Version   string
//	CreatedAt string
//	Hashrate  string
//	Watt      string
//	TopWatt   string
//	WuhaiWatt string
//	CoinType  string
//	HouqiWatt string
//}
type GoodsNameO struct {
	GoodsName
	CreatedAtFormat string
}


func init()  {
	orm.RegisterModel(new(GoodsName),new(GoodsNameO))
}

// 通过id查找用户
func (goodsName *GoodsName) GetById() (*GoodsName, error) {
	o := orm.NewOrm()
	err := o.Read(goodsName, "id")
	return goodsName, err
}

//获取列表
func GetList(pageNo int,pageSize int,keyword string,typeId string) ([]*GoodsName,int)  {
	//高级查询
	//var goodsList []*GoodsName
	//offset := (pageNo-1)*pageSize
	//query := orm.NewOrm().QueryTable("goods_name")
	//query = query.Filter("type_id",typeId)
	//total, _ := query.Count()
	//totalCount := int(total)
	//query.OrderBy("-created_at").Offset(offset).Limit(pageSize).All(&goodsList)
	//return  goodsList,totalCount
	//原生sql查询
	o := orm.NewOrm()
	var goodsList []*GoodsName
	offset := (pageNo-1)*pageSize
	var sql = fmt.Sprintf("SELECT t.* FROM goods_name t")
	if typeId != "" {
		sql += fmt.Sprintf(" where t.type_id =%s", typeId)
	}
	totalCount,_ := o.Raw(sql).QueryRows(&goodsList)
	sql += fmt.Sprintf(" limit %d,%d", offset,pageSize)
	o.Raw(sql).QueryRows(&goodsList)
	return goodsList,int(totalCount)
	//构造查询
	//o := orm.NewOrm()
	//var goodsList []*GoodsName
	//offset := (pageNo-1)*pageSize
	//qb, _ := orm.NewQueryBuilder("mysql")
	//qb = qb.Select("*").From("goods_name")
	//if typeId != "" {
	//	qb.Where("type_id = ?")
	//}
	//sql := qb.String()
	//totalCount,_ := o.Raw(sql, typeId).QueryRows(&goodsList)
	//qb.Limit(pageSize).Offset(offset)
	//sqla := qb.String()
	//o.Raw(sqla, typeId).QueryRows(&goodsList)
	//return goodsList,int(totalCount)
}

