package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GoHold/controllers/subscribe:SubScribeController"] = append(beego.GlobalControllerRouter["GoHold/controllers/subscribe:SubScribeController"],
		beego.ControllerComments{
			Method: "Ssr",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/subscribe:SubScribeController"] = append(beego.GlobalControllerRouter["GoHold/controllers/subscribe:SubScribeController"],
		beego.ControllerComments{
			Method: "SubList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
