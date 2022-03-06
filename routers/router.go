// @APIVersion 1.0.0
// @Title ssoer
// @Description SingleSignOn application
package routers

import (
	"ssoer/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/ping",
			beego.NSInclude(
				&controllers.PingController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
