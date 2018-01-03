package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"] = append(beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"] = append(beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"] = append(beego.GlobalControllerRouter["GoHold/controllers/mine:UserController"],
		beego.ControllerComments{
			Method: "UpLoadUserPhoto",
			Router: `/uploadPhoto`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

}
