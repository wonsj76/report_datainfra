package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SReportsController struct {
	beego.Controller
}

type InServerData struct {
	SCountry   string `json:"scoun"`
	SBinary    string `json:"sbin"`
	SVersion   string `json:"sver"`
	SPublicIP  string `json:"spu"`
	SPraviteIP string `json:"spr"`
	SFarmIP    string `json:"sfar"`
}

func (k *SReportsController) GetServer() {
	//test
	var inserverdata InServerData

	logs.NewLogger(3000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/server_project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true}`)

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inserverdata)

	if err != nil {
		logs.Info("Server Error : ", inserverdata)
	} else {
		logs.Info("Server : ", inserverdata)
	}

	defer k.ServeJSON()
}
