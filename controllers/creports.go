package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CReportsController struct {
	beego.Controller
}

type InClientData struct {
	ClCountry  string `json:"clcoun"`
	ClVersion  string `json:"clver"`
	ClServerIP string `json:"clsip"`
	ClSVersion string `json:"clsver"`
	ClFile     string `json:"clfile"`
	ClConnect  string `json:"clcon"`
}

func (b *CReportsController) GetClient() {

	var inclientdata InClientData

	logs.NewLogger(3000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/server_project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true}`)

	body, _ := ioutil.ReadAll(b.Ctx.Request.Body)
	err := json.Unmarshal(body, &inclientdata)

	if err != nil {
		logs.Info("Client Error : ", inclientdata)
	} else {
		logs.Info("Client : ", inclientdata)
	}

	defer b.ServeJSON()
}
