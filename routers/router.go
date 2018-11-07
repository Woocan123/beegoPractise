package routers

import (
	"beegoPractise/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/getById", &controllers.GoodsNameController{},"GET:GetById")
	beego.Router("/getList", &controllers.GoodsNameController{},"GET:GetList")

	beego.Router("/add", &controllers.GoodsNameController{},"PUT:AddOrUpdate")
	beego.Router("/delete", &controllers.GoodsNameController{},"DELETE:Delete")
}
