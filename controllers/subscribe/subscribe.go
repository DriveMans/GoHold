package subscribe

import (
	"github.com/astaxie/beego"
	"GoHold/models"
	"GoHold/controllers/tool"
	"strconv"
)

type SubScribeController struct {
	beego.Controller
}

// @Title 获取订阅列表
// @Description 订阅列表
// @Param	token		query 	string	true		"令牌"
// @Success 200 {string} 响应成功!
// @Failure 响应失败
// @router /list [post]
func (self *SubScribeController)SubList(){
	token := self.Input().Get("token")

	user,err := models.TokenVerifyUser(token)

	if err != nil {
		self.Data["json"] = tool.ResultData{Code:999,Message:"token 失效，请重新登录"}
		self.ServeJSON()
		return
	}

	ssrs , err := models.SubAllData(user.Id)
	if err != nil {
		self.Data["json"] = tool.ResultData{Code:0,Message:err.Error()}
		self.ServeJSON()
		return
	}

	var resultArr []interface{}

	for _, value := range ssrs {

		//获取用户信息
		beUser,err := models.SearchUserWithUserId(value.SubjectId)
		if err != nil {
			beego.Error(err)
			continue
		}

		m := make(map[string]interface{})
		m["Account"] = beUser.Account
		m["ArticleNum"] = beUser.ArticleNum
		m["Contribute"] = beUser.Contribute
		m["UserId"] = beUser.Id
		m["NikeName"] = beUser.NikeName

		photoPath,err := models.GetPictureWithId(beUser.Photo)
		if err != nil {
			m["Photo"] = ""
		}else {
			m["Photo"] = photoPath
		}


		m["RssNum"] = beUser.RssNum

		item := tool.StructToMap(*value)
		delete(item,"Id")
		delete(item,"CreateTime")
		delete(item,"UpdateTime")
		delete(item,"Status")
		delete(item,"SubjectId")
		delete(item,"ObserverId")

		item["Ids"] = value.Id

		item["user_info"] = m

		resultArr = append(resultArr,item)
	}

	self.Data["json"] = tool.ResultData{Code:200,Message:"获取成功",Data:resultArr}
	self.ServeJSON()
}

// @Title 订阅某个用户
// @Description 订阅
// @Param	token		query 	string	true		"令牌"
// @Param	be_user_id	query	string  true		"被订阅人的ID"
// @Success 200 {string} 响应成功!
// @Failure 响应失败
// @router / [post]
func (self *SubScribeController)Ssr(){
	token := self.Input().Get("token")
	be_userid := self.Input().Get("be_user_id")
	be_userid_int,_ := strconv.ParseInt(be_userid,10,64)
	user,err := models.TokenVerifyUser(token)
	if err != nil {
		self.Data["json"] =tool.ResultData{Code:999,Message:tool.ReturnErrorTokenFailure}
		self.ServeJSON()
		return
	}
	
	//根据user_id查找用户
	be_user , err := models.SearchUserWithUserId(be_userid_int)
	if err != nil {
		self.Data["json"] = tool.ResultData{Code:0,Message:tool.ReturnErrorUserNonExistence}
		return
	}

	err = models.SubScribe(be_user.Id,user.Id)
	if err != nil {
		self.Data["json"] = tool.ResultData{Code:0,Message:err.Error()}
		self.ServeJSON()
		return
	}

	self.Data["json"] = tool.ResultData{Code:200,Message:tool.ReturnOperateSuccess}
	self.ServeJSON()

	
}