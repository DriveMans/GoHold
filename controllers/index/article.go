package index

import (
	"GoHold/controllers/tool"
	"GoHold/models"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
)

type CommentResult struct {
	Comment 	*models.GhComment/*评论主体*/
	NikeName 	string/*昵称*/
	Photo 		string/*头像*/
	Reply 		string/*回复内容*/
	IsReply 	string/*主人是否回复 1 true 0 false*/
}

type ArticleController struct {
	beego.Controller
}

// @Title 获取首页数据
// @Description 获取首页数据
// @Param	index_page		query 	string	true		"页码"
// @Success 200 {string} 获取成功!
// @Failure 获取失败
// @router / [get,post]
func (self *ArticleController) Get() {

	indexPage := self.Input().Get("index_page")

	articles, err := models.GetIndexDataWithIndexPage(indexPage)

	if err != nil {
		beego.Error(err)

		//这里有点多余   不过也还好啦
		self.Data["json"] = tool.ResultData{Code: 0, Message: "没有更多数据啦！"}
		self.ServeJSON()
		return
	}

	if len(articles) > 1 {
		self.Data["json"] = tool.ResultData{Code: 200, Message: "获取成功", Data: articles}
	} else {
		self.Data["json"] = tool.ResultData{Code: 0, Message: "没有更多数据啦！"}
	}

	self.ServeJSON()

}

// @Title 浏览文章（隐式操作）
// @Description 浏览文章
// @Param	aid		query 	string	true		"文章id"
// @Success 200 {string} 响应成功!
// @Failure 响应失败
// @router /view [get,post]
func (self *ArticleController) ViewArticle() {

	aid := self.Input().Get("aid")
	token := self.Input().Get("token")

	models.ViewArticle(aid,token)

	self.Data["json"] = tool.ResultData{Code:200,Message:"success"}
	self.ServeJSON()

}

// @Title 获取banner数据
// @Description 获取banner数据
// @Param	aid		query 	string	true		"文章id"
// @Success 200 {string} 响应成功!
// @Failure 响应失败
// @router /getBanner [get,post]
func (self *ArticleController) GetBanner() {
	self.Data["json"] = tool.ResultData{Code:200,Message:"获取banner成功"}
	self.ServeJSON()

}


// @Title 用户发布文章
// @用户发布文章
// @Param	token		query 	string	true		"口令"
// @Param	title		query 	string	true		"标题"
// @Param	url			query 	string	true		"文章引用的url"
// @Param	img_url		query 	string	true		"文章图片路径"
// @Param	tags		query 	string	true		"标签组"
// @Success 200 {string} 发布成功!
// @Failure 发布失败
// @router /publishArticle [get,post]
func (this *ArticleController) PublishArticle() {
	token := this.Input().Get("token")
	title := this.Input().Get("title")
	url := this.Input().Get("url")
	imageUrl := this.Input().Get("img_url")
	tags := this.Input().Get("tags")

	//tags 和 imageUrl 可以为空 用户可以不添加

	if len(token) < 1 || len(title) < 1 || len(url) < 1 {
		this.Data["json"] = tool.ResultData{Code: 0, Message: "参数不完整"}
		this.ServeJSON()
		return
	}

	user, err := models.TokenVerifyUser(token)

	if err != nil {
		beego.Error(err)
		this.Data["json"] = tool.ResultData{Code: 999, Message: "Token 失效请重新登录"}
		this.ServeJSON()
		return
	} else {
		article := new(models.GhArticle)
		article.Title = title
		article.Url = url
		article.Img = imageUrl
		article.Tags = tags
		article.UserId = user.Id
		article.Status = 1

		err = models.InsertArticle(article)
		if err != nil {
			this.Data["json"] = tool.ResultData{Code: 0, Message: err.Error()}
		} else {
			this.Data["json"] = tool.ResultData{Code: 200, Message: "文章发布成功!"}
		}
		this.ServeJSON()
	}
}


// @Title 评论文章
// @Description 用户评论文章
// @Param	token			query 	string	true		"口令"
// @Param	content			query 	string	true		"内容"
// @Param	articleId		query 	string	true		"文章id"
// @Param	commentId		query 	string	true		"评论id"
// @Param	beUserId		query 	string	true		"被评论人id"
// @Success 200 {string} 评论成功!
// @Failure 评论失败
// @router /replyArticle [get,post]
func (self *ArticleController) ReplyArticle() {

	token := self.Input().Get("token")
	content := self.Input().Get("content")
	aid := self.Input().Get("articleId")
	beUserId := self.Input().Get("beUserId")
	beCommentId := self.Input().Get("beCommentId")

	if len(token) < 1 || len(content) < 1 || len(aid) < 1 {
		self.Data["json"] = tool.ResultData{Code:0,Message:"参数不完整"}
		self.ServeJSON()
		return
	}


	if len(content) > 200 {
		self.Data["json"] = tool.ResultData{Code:0,Message:"评论内容过长 "}
		self.ServeJSON()
		return
	}
	uid ,err :=  models.TokenExchangeUserId(token)

	if err != nil {
		beego.Error(err)
		self.Data["json"] = tool.ResultData{Code:999,Message:"token已失效，请重新登录"}
		self.ServeJSON()
		return
	}

	comment := new(models.GhComment)
	comment.UserId = uid
	comment.BeUserId,_ = strconv.ParseInt(beUserId,10,64)
	comment.ArticleId,_ = strconv.ParseInt(aid,10,64)
	//comment.BeCommentId,_ = strconv.ParseInt(cid,10,64)
	comment.Content = content
	comment.Status = 1
	comment.Tree = 1

	if len(beUserId) > 0 {

		if len(beCommentId) < 1 {
			self.Data["json"] = tool.ResultData{Code:1,Message:"缺少绑定参数"}
			self.ServeJSON()
			return
		}
		comment.Tree = 2
		comment.BeCommentId,_ = strconv.ParseInt(beCommentId,10,64)
		comment.BeUserId,_ = strconv.ParseInt(beUserId,10,64)
	}

	if len(beCommentId) > 0 {
		if len(beUserId) < 1 {
			self.Data["json"] = tool.ResultData{Code:0,Message:"缺少绑定参数"}
			self.ServeJSON()
			return
		}
		comment.Tree = 2
		comment.BeCommentId,_ = strconv.ParseInt(beCommentId,10,64)
		comment.BeUserId,_ = strconv.ParseInt(beUserId,10,64)
	}

	err = models.ReplyArticle(comment)

	if err != nil {
		beego.Error(err)
		self.Data["json"] = tool.ResultData{Code:0,Message:"回复失败"}
	}else {
		self.Data["json"] = tool.ResultData{Code:200,Message:"回复成功"}
	}
	self.ServeJSON()

}

// @Title 文章点赞
// @Description 文章点赞
// @Param	token			query 	string	true		"口令"
// @Param	articleId		query 	string	true		"文章id"
// @Success 200 {string} 点赞成功!
// @Failure 点赞失败
// @router /likeArticle [post]
func (self *ArticleController) LikeArticle () {
	token := self.Input().Get("token")
	aid := self.Input().Get("articleId")

	if len(token) < 1 || len(aid) < 1 {
		self.Data["json"] = tool.ResultData{Code:0,Message:tool.ReturnErrorParameNotNull}
		self.ServeJSON()
		return
	}

	uid , err := models.TokenExchangeUserId(token)

	if err != nil {
		beego.Error(err)
		self.Data["json"]= tool.ResultData{Code:999,Message:tool.ReturnErrorTokenFailure}
		self.ServeJSON()
		return
	}
	fmt.Println("----------a")
	err = models.LikeArticle(uid,aid)
	fmt.Println("----------b")
	if err != nil {
		beego.Error(err)
		self.Data["json"]= tool.ResultData{Code:0,Message:err.Error()}
		self.ServeJSON()
		return
	}

	self.Data["json"]= tool.ResultData{Code:200,Message:"操作成功"}
	self.ServeJSON()

}

// @Title 文章收藏
// @Description 用户收藏  和点赞差不多
// @Param	token			query 	string	true		"口令"
// @Param	articleId		query 	string	true		"文章id"
// @Success 200 {string} 点赞成功!
// @Failure 点赞失败
// @router /collectArticle [post]
func (self *ArticleController) CollectArticle() {
	token := self.Input().Get("token")
	aid := self.Input().Get("articleId")

	if len(token) < 1 || len(aid) < 1 {
		self.Data["json"] = tool.ResultData{Code:0,Message:tool.ReturnErrorParameNotNull}
		self.ServeJSON()
	}

	uid , err := models.TokenExchangeUserId(token)

	if err != nil {
		beego.Error(err)
		self.Data["json"]= tool.ResultData{Code:999,Message:tool.ReturnErrorTokenFailure}
		self.ServeJSON()
	}

	fmt.Println("----------c")
	err = models.CollectArticle(uid,aid)
	fmt.Println("----------d")
	if err != nil {
		beego.Error(err)
		self.Data["json"]= tool.ResultData{Code:0,Message:err.Error()}
		self.ServeJSON()
	}

	self.Data["json"]= tool.ResultData{Code:200,Message:"操作成功"}
	self.ServeJSON()
}


// @Title 查看此文章用户是否点赞
// @Description 查看此文章用户是否点赞
// @Param	token			query 	string	true		"口令"
// @Param	articleId		query 	string	true		"文章id"
// @Success 200 {string} 点赞成功!
// @Failure 点赞失败
// @router /checkLike [post]
func (self *ArticleController)CheckLike()  {
	token := self.Input().Get("token")
	aid := self.Input().Get("articleId")

	if len(token) < 1 || len(aid) < 1 {
		self.Data["json"] = tool.ResultData{Code:0,Message:tool.ReturnErrorParameNotNull}
		self.ServeJSON()
	}

	uid , err := models.TokenExchangeUserId(token)

	if err != nil {
		beego.Error(err)
		self.Data["json"]= tool.ResultData{Code:999,Message:tool.ReturnErrorTokenFailure}
		self.ServeJSON()
	}

	if models.CheckLike(uid,aid) {
		self.Data["json"] = tool.ResultData{Code:200,Message:"已点赞",Data:map[string]string{"status":"1"}}
	}else {
		self.Data["json"] = tool.ResultData{Code:200,Message:"未点赞",Data:map[string]string{"status":"0"}}
	}

	self.ServeJSON()

}

// @Title 获取文章的全部评论 	Not Token
// @Description 获取文章的全部评论
// @Param	articleId		query 	string	true		"文章id"
// @Success 200 {string} 点赞成功!
// @Failure 点赞失败
// @router /allComment [post]
func (self *ArticleController)AllComment()  {
	articleId := self.Input().Get("articleId")

	if len(articleId) < 1 {
		self.Data["json"] = tool.ResultData{Code:0,Message:tool.ReturnErrorParameNotNull}
		self.ServeJSON()
		return
	}

	comments,err := models.GetAllComment(articleId)
	if err != nil {
		fmt.Printf("id:%s 没有值",articleId)
		self.Data["json"] = tool.ResultData{Code:1,Message:err.Error()}
		self.ServeJSON()
		return
	}

	var sources []*CommentResult

	for _,comment := range comments {


		comres := new(CommentResult)
		comres.Comment = comment

		user,err := models.SearchUserWithUserId(comment.UserId)
		if err != nil {

			self.Data["json"] = tool.ResultData{Code:0,Message:"当前没有评论"}
			self.ServeJSON()
			return
		}else {
			//m := make(map[string]string)
			//UserId := strconv.FormatInt(user.Id,10)
			//m["UserId"] = UserId
			//m["NikeName"] = user.NikeName
			//m["Photo"] = ""
			comres.Photo = ""
			comres.NikeName = user.NikeName
			comres.IsReply = "0"
			comres.Reply = ""

			picture_path,err := models.GetPictureWithId(user.Photo)
			if err == nil {
				//host := beego.AppConfig.String("httphost")
				//prot := beego.AppConfig.String("httpport")
				//
				//hostPath := "http://" + host + ":" + prot + picture_path
				comres.Photo = picture_path

			}

			sources = append(sources,comres)
			//用User 找头

		}
	}


	self.Data["json"] = tool.ResultData{Code:200,Message:"获取成功",Data:sources}
	self.ServeJSON()
}
