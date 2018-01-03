package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "AllComment",
			Router: `/allComment`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "CheckLike",
			Router: `/checkLike`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "CollectArticle",
			Router: `/collectArticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "GetBanner",
			Router: `/getBanner`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "LikeArticle",
			Router: `/likeArticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "PublishArticle",
			Router: `/publishArticle`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "ReplyArticle",
			Router: `/replyArticle`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"] = append(beego.GlobalControllerRouter["GoHold/controllers/index:ArticleController"],
		beego.ControllerComments{
			Method: "ViewArticle",
			Router: `/view`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

}
