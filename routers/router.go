package routers

import (
	"suanli/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/getById", &controllers.GoodsNameController{},"GET:GetById")
	beego.Router("/getList", &controllers.GoodsNameController{},"GET:GetList")
}
