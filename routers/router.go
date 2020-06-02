package routers

import (
	"github.com/astaxie/beego"
	"github.com/ndcinfra/report_datainfra/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/creports",
			beego.NSRouter("/getclient", &controllers.CReportsController{}, "post:GetClient"),
		),
		beego.NSNamespace("/sreports",
			beego.NSRouter("/getserver", &controllers.SReportsController{}, "post:GetServer"),
		),
	)
	beego.AddNamespace(ns)
}
