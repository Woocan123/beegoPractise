package routers

import (
	"beegoPractise/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/goodsName/getById", &controllers.GoodsNameController{},"GET:GetById")
	beego.Router("/goodsName/getList", &controllers.GoodsNameController{},"GET:GetList")

	beego.Router("/goodsName/add", &controllers.GoodsNameController{},"PUT:AddOrUpdate")
	beego.Router("/goodsName/delete", &controllers.GoodsNameController{},"DELETE:Delete")

	beego.Router("/electric/add", &controllers.ElectricMeterController{},"PUT:AddOrUpdate")
	beego.Router("/electric/getById", &controllers.ElectricMeterController{},"GET:GetById")
}
