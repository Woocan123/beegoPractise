package controllers

import "github.com/astaxie/beego"

type RepertoryController struct {
	beego.Controller
}

func (r *RepertoryController) consumerMachineList()  {
	 bdcId :=r.GetString("bdcId")
	for a :=0; a<3 ; a++  {
		if a==0 {

		}else if a==1 {

		}
	}
}
