package controllers

import (
	. "beegoPractise/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type RepertoryController struct {
	beego.Controller
}

func (r *RepertoryController) ConsumerMachineList()  {
	bdcId :=r.GetString("bdcId")
	for a :=0; a<3 ; a++  {
		if a==0 {
			consumer := GetConsumers(strconv.Itoa(a),bdcId)
			for _,value := range consumer{
				miner := GetMiners(strconv.Itoa(a),bdcId,value.Id)
				for _,minerValue := range miner{
					machineGoods := GetMachineGoods(strconv.Itoa(a),bdcId,value.Id,minerValue.Mine)
					minerValue.ConsumerGoods = machineGoods
					fmt.Print(machineGoods)
				}

			}

		}else if a==1 {

		}
	}
}
