package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
)

type GhUser struct {
	Id         int64     //主键
	Account    string    //账户
	Password   string    //账号
	Token      string    //令牌
	NikeName   string	 `orm:"default()"`
	Contribute int64     `orm:"default(0)"`                  //贡献值
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	LoginTime  time.Time `orm:"null"`                        //登录时间
	LoginIpv4  string    //登录的IP地址
	UseVersion string    //当前使用APP版本
	PushId     string    //推送ID
	Status     int64     `orm:"default(1)"` //状态 0:被封禁用户  1:正常用户   2:管理员   3:超级管理员
	ArticleNum int64     `orm:"default(0)"` //文章数
	RssNum     int64     `orm:"default(0)"` //订阅数
	Photo      int64	 `orm:"null"`
	RegisterIp string  	 `orm:"null"`

}

/*
	创建用户
*/
func CreateUser(account, password ,token ,registerip,photoid string) error {
	user := new(GhUser)
	user.Account = account
	user.Password = password
	user.Token = token
	user.Status = 1
	user.RegisterIp = registerip
	pid ,_ := strconv.ParseInt(photoid,10,64)
	user.Photo = pid

	_, err := db.Insert(user)
	return err
}

/*
	账户名取用户（手机账号）
*/
func QueryUserWithAccount(account string) (num int64, err error) {
	num, err = db.QueryTable("gh_user").Filter("account", account).Count()
	return num, err
}

/*
	用户登录
*/
func UserLogin(account, password string) (*GhUser, error) {
	user := new(GhUser)

	//选查询手机号

	err := db.QueryTable("gh_user").Filter("account",account).One(user)
	if err != nil {
		return nil,fmt.Errorf("账户不存在")
	}

	if password != user.Password {
		return nil,fmt.Errorf("密码不正确")
	}

	return user,nil
	//
	//err := db.QueryTable("gh_user").Filter("account", account).Filter("password", password).One(user)
	//return user, err
}

/*
	更新用户Token
*/
func UserUpdateToken(uid int64, token ,pushId string) error {
	_, err := db.QueryTable("gh_user").Filter("id", uid).Update(orm.Params{"token": token, "login_time": time.Now(),"push_id":pushId})
	return err
}

/*
	Token 验证 User
*/
func TokenVerifyUser(token string) (*GhUser, error) {
	user := new(GhUser)
	err := db.QueryTable("gh_user").Filter("token", token).One(user)
	return user, err
}

/*
	Token 换 UserID
*/
func TokenExchangeUserId(token string) (uid int64,err error){

	var user = new(GhUser)

	err = db.QueryTable("gh_user").Filter("token",token).One(user,"id")
	return  user.Id,err
}

/*
	根据ID查找用户
*/
func SearchUserWithUserId(userid int64)(*GhUser,error)  {
	//user := GhUser{Id:userid}
	var user GhUser

	err := db.QueryTable("gh_user").Filter("id",userid).One(&user,"NikeName","Id","Photo")
	if err != nil {
		beego.Error(err)
		return nil,err
	}else {
		return &user,nil
	}

}