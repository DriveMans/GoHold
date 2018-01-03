package mine

import (
	"github.com/astaxie/beego"
	"GoHold/controllers/tool"
	"GoHold/models"
	"GoHold/libs"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"time"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"github.com/1l0/identicon"
)

// 用户行为
type UserController struct {
	beego.Controller
}

// @Title User what fuck
// @Description 用来登陆的接口
// @Param	account		body 	string	true		"账户名（手机号）"
// @Param	password	body 	string	true		"密码（6-12位）"
// @Success 200 {string} 登录成功!
// @Failure 账户名或密码不对
// @router /login [post]
func (self *UserController) Login() {

	account := self.Input().Get("account")
	password := self.Input().Get("password")
	pushId := self.Input().Get("push_id")

	if len(pushId) < 1 {
		pushId = ""
	}

	if len(account) < 1 || len(password) < 1 {
		self.Data["json"] = tool.ResultData{Code: 0, Message: "用户名或密码不正确1"}
		self.ServeJSON()
		return
	}

	user, err := models.UserLogin(account, password)

	if err != nil {
		self.Data["json"] = tool.ResultData{Code: 0, Message: err.Error()}
		self.ServeJSON()
		return
	}

	token := libs.GetRandomString(50)
	err = models.UserUpdateToken(user.Id, token,pushId)
	if err != nil {
		self.Data["json"] = tool.ResultData{Code: 0, Message: err.Error()}
	} else {
		self.Data["json"] = tool.ResultData{Code: 200, Message: "登录成功", Data: map[string]string{"token": token}}
	}

	self.ServeJSON()
}

// @Title 用户注册接口
// @Description 用户注册
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} 登录成功!
// @Failure 账户名或密码不对
// @router /register [get,post]
func (this *UserController) Register() {
	account := this.Input().Get("account")
	password := this.Input().Get("password")

	if len(account) < 1 {
		this.Data["json"] = tool.ResultData{Code: 200, Message: "请输入账户"}
	} else if len(password) < 1 {
		this.Data["json"] = tool.ResultData{Code: 200, Message: "请输入密码"}
	} else {
		if len(password) >= 6 && len(password) <= 16 {
			num, _ := models.QueryUserWithAccount(account)
			if num > 0 {
				this.Data["json"] = tool.ResultData{Code: 0, Message: "此账户已注册"}
			} else {
				id:=identicon.New()
				id.Type = identicon.Normal
				id.Theme = identicon.White

				//创建文件夹
				times := time.Now()
				dir := times.Format("2006-01-02")
				path  := "/static/images/systemPhoto/" + dir +"/"

				fileName := path + account

				if libs.CheckFileIsExist("." + path) {
					id.GeneratePNGToFile( "." + fileName )
				}else {
					err := os.MkdirAll("." + path,0777)
					if err != nil {
						this.Data["json"] = tool.ResultData{Code:3,Message:err.Error()}
						this.ServeJSON()
					}else {
						id.GeneratePNGToFile("." + fileName )
					}
				}

				host := beego.AppConfig.String("httphost")
				port := beego.AppConfig.String("httpport")

				hostPath := "http://" + host + ":" + port + fileName + ".png"
				fmt.Println(hostPath)

				picture := new(models.GhPicture)
				picture.Remark = "用户默认头像"
				picture.Path = fileName + ".png"
				picture.Suffix = "png"

				pid ,ok := models.PictureInster(picture)

				if !ok {
					this.Data["json"] = tool.ResultData{Code:0,Message:"insert picture error"}
					this.ServeJSON()
				}else {

					addr := this.Ctx.Request.RemoteAddr
					token := libs.GetRandomString(50)
					err := models.CreateUser(account, password,token,addr,strconv.FormatInt(pid,10))
					if err != nil {
						this.Data["json"] = tool.ResultData{Code: 0, Message: err.Error()}
					} else {
						this.Data["json"] = tool.ResultData{Code: 200, Message: "注册成功",Data:map[string]string{"token":token,"user_photo":hostPath}}
					}
				}
			}
		} else {
			this.Data["json"] = tool.ResultData{Code: 200, Message: "密码不符合规范"}
		}
	}
	this.ServeJSON()
}

// @Title 用户上传头像
// @Description 上传二进制流图片   返回url
// @Param	photo		query 	string	true		"字符流"
// @Success 200 {string} 上传成功!
// @Failure 上传失败
// @router /uploadPhoto [get,post]
func (self *UserController) UpLoadUserPhoto() {
	str := self.Input().Get("img")
	times := time.Now()
	dir := times.Format("2006-01-02")
	path  := "/static/images/userPhoto/" + dir +"/"
	ddd, _ := base64.StdEncoding.DecodeString(str) //成图片文件并把文件写入到buffer
	//err := ioutil.WriteFile("./output.jpg", ddd, 0666)

	r := libs.GetRandomString(32)
	filename := r + ".jpg"

	if libs.CheckFileIsExist("." + path) {
		fmt.Println("路径存在")
		err := ioutil.WriteFile("." + path + filename,ddd,0666)
		if err != nil {
			self.Data["json"] = tool.ResultData{Code:0,Message:err.Error()}
			self.ServeJSON()
		}else {
			host := beego.AppConfig.String("httphost")
			prot := beego.AppConfig.String("httpport")

			hostPath := "http://" + host + ":" + prot + path + filename
			fmt.Println(hostPath)

			picture := new(models.GhPicture)
			picture.Remark = "用户头像"
			picture.Path = path + filename
			picture.Suffix = "jpg"

			pid ,ok := models.PictureInster(picture)

			if !ok {
				self.Data["json"] = tool.ResultData{Code:0,Message:"插入图片出错"}
				self.ServeJSON()
			}else {
				self.Data["json"] = tool.ResultData{Code:200,Message:"上传成功",Data:map[string]string{"imgUrl":hostPath,"picutre_id":strconv.FormatInt(pid,10)}}
				self.ServeJSON()
			}
		}
	}else {
		err := os.MkdirAll("." + path,0777)

		if err != nil {
			self.Data["json"] = tool.ResultData{Code:3,Message:err.Error()}
			self.ServeJSON()
		}

		err = ioutil.WriteFile("." + path + filename,ddd,0666)

		if err != nil {
			self.Data["json"] = tool.ResultData{Code:1,Message:err.Error()}
			self.ServeJSON()
		}else {
			host := beego.AppConfig.String("httphost")
			prot := beego.AppConfig.String("httpport")

			hostPath := "http://" + host + ":" + prot + path + filename
			fmt.Println(hostPath)

			picture := new(models.GhPicture)
			picture.Remark = "用户头像"
			picture.Path = path + filename
			picture.Suffix = "jpg"

			pid ,ok := models.PictureInster(picture)

			if !ok {
				self.Data["json"] = tool.ResultData{Code:0,Message:"插入图片出错"}
				self.ServeJSON()
			}else {
				self.Data["json"] = tool.ResultData{Code:200,Message:"上传成功",Data:map[string]string{"imgUrl":hostPath,"picutre_id":strconv.FormatInt(pid,10)}}
				self.ServeJSON()
			}
		}
	}
}




func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}