// @APIVersion 1.0.0
// @Title GoHold API文档
// @Description 每个API接口所需要的参数及其请求方式都在文档中
// @Contact 共电科技
// @TermsOfServiceUrl http://www.ipower001.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GoHold/controllers/index"
	"github.com/astaxie/beego"
	"GoHold/controllers/mine"
	"GoHold/controllers/subscribe"
)

func init() {

	  ns :=	beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&mine.UserController{},
				),
			),
		  	beego.NSNamespace("/ssr",
			  	beego.NSInclude(
				  	&subscribe.SubScribeController{},
			  	),
		 	 ),
			beego.NSNamespace("/index",
				beego.NSInclude(
					&index.ArticleController{},
				),
			),
		)


	//
	//web_ns := beego.NewNamespace("/app",
	//	beego.NSNamespace("/v1",
	//		beego.NSNamespace("/user",
	//			beego.NSInclude(
	//				&mine.UserController{},
	//			),
	//		),
	//		beego.NSNamespace("/index",
	//			beego.NSInclude(
	//				&index.ArticleController{},
	//			),
	//		),
	//	),
	//)


	beego.SetStaticPath("/static/userPhoto","static")

	beego.AddNamespace(ns)
}
