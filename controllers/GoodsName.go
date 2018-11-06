package controllers

import (
	"github.com/astaxie/beego"
	. "suanli/models"
	)


type GoodsNameController struct {
	beego.Controller
}

func (g *GoodsNameController) GetById()  {
	id := g.GetString("id")
	var goodsName = new(GoodsName)
	goodsName.Id = id
	data,_ := goodsName.GetById()
	var result = new(GoodsNameO)
	result.GoodsName = *data
	result.CreatedAtFormat = data.CreatedAt.Format("2006-01-02 15:04:05")
	g.Data["json"] = map[string]interface{}{"data":result}
	g.ServeJSON()
}

func (g *GoodsNameController) GetList()  {
	keyword := g.GetString("keyword")
	pageNo, _ := g.GetInt("pageNo")
	pageSize, _ := g.GetInt("pageSize")
	typeId := g.GetString("typeId")
	result,total := GetList(pageNo,pageSize,keyword,typeId)
	g.Data["json"] = map[string]interface{}{"data":result,"total":total}
	g.ServeJSON()
}
