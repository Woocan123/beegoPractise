package controllers

import (
	. "beegoPractise/models"
	"github.com/astaxie/beego"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)


type GoodsNameController struct {
	beego.Controller
}

func (g *GoodsNameController) GetById()  {
	id := g.GetString("id")
	var goodsName = new(GoodsName)
	goodsName.Id = id
	data,_ := goodsName.GetById()
	var result GoodsNameO
	result.GoodsName = *data
	result.CreatedAtFormat = data.CreatedAt.Format("2006-01-02 15:04:05")
	g.Data["json"] = map[string]interface{}{"msg":"成功","data":result,"code":0}
	g.ServeJSON()
}

func (g *GoodsNameController) GetList()  {
	keyword := g.GetString("keyword")
	pageNo, _ := g.GetInt("pageNo")
	pageSize, _ := g.GetInt("pageSize")
	typeId := g.GetString("typeId")
	result,total := GetList(pageNo,pageSize,keyword,typeId)
	 //newResult := list.New()
	 //添加操作一般用slice，list则在删除操作中用
	 //var a []*GoodsNameO
	for _,v := range result{
		v.CreatedAtFormat = v.CreatedAt.Format("2006-01-02 15:04:05")
		//var goodsNameo = new(GoodsNameO)
		//goodsNameo.GoodsName = *v
		//goodsNameo.CreatedAtFormat = v.CreatedAt.Format("2006-01-02 15:04:05")
		//a = append(a, goodsNameo)
	}
	
	data := make(map[string]interface{})
	data["list"] = result
	data["total"] = total
	g.Data["json"] = map[string]interface{}{"msg":"成功","data":data,"code":0}
	g.ServeJSON()
}

func (g *GoodsNameController) AddOrUpdate()  {
	id := g.GetString("id")
	if id =="" {
		id =uuid.Must(uuid.NewV4()).String()
		id = strings.Replace(id,"-","",-1)
	}
	var goodsName GoodsName
	if err := g.ParseForm(&goodsName); err != nil {

	}
	if id =="" {
		goodsName.CreatedAt = time.Now()
	}
	goodsName.Id = id
	goodsName.AddOrUpdate()
	g.Data["json"] = map[string]interface{}{"msg":"成功","code":0}
	g.ServeJSON()
}

func (g *GoodsNameController) Delete()  {
	id := g.GetString("id")
	var goodsName GoodsName
	goodsName.Id = id
	num := goodsName.Delete()
	if num != 0{
		g.Data["json"] = map[string]interface{}{"msg":"成功","code":0}
	}else {
		g.Data["json"] = map[string]interface{}{"msg":"失败","code":1}
	}
	g.ServeJSON()
}
