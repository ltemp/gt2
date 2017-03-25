package routers

import (
	"gett2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/driver",
			beego.NSInclude(
				&controllers.DriverController{},
			),
		),
		beego.NSNamespace("/metric",
			beego.NSInclude(
				&controllers.MetricController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
