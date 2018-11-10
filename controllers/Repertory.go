package controllers

import (
	. "beegoPractise/models"
	"github.com/astaxie/beego"
	"strconv"
)

type RepertoryController struct {
	beego.Controller
}

func (r *RepertoryController) ConsumerMachineList()  {
	bdcId :=r.GetString("bdcId")
	result := make(map[string]interface{})
	for a :=0; a<3 ; a++  {
		if a==0 {
			consumer := GetConsumers(strconv.Itoa(a),bdcId)
			for _,value := range consumer{
				miner := GetMiners(strconv.Itoa(a),bdcId,value.Id)
				for _,minerValue := range miner{
					machineGoods := GetMachineGoods(strconv.Itoa(a),bdcId,value.Id,minerValue.Mine)
					minerValue.ConsumerGoods = machineGoods
				}
				value.MinerO = miner
			}
			result["machines"] = consumer
		}else if a==1 {

		}
	}
	r.Data["json"] = map[string]interface{}{"msg":"成功","data":result,"code":0}
	r.ServeJSON()
}
