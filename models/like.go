package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
)

/*
	点赞
*/
type GhLike struct {
	Id 			int64			//主键
	Status  	int64			//状态
	UserId   	int64			//用户id
	ArticleId	int64			//文章id
	Tags 		string			//标签组
	CreateTime 	time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}

/*
	文章点赞
*/
func LikeArticle(uid int64,aid string)error  {

	int_aid,_ := strconv.ParseInt(aid,10,64)

	//如果已经点赞  操作就是取消赞
	//如果没有点赞  操作就是点赞
	if CheckLike(uid,aid) {
		fmt.Println("已经点赞了，开始取消赞")
		//已经点赞  取消赞
		_,err := db.QueryTable("gh_like").Filter("user_id",uid).Filter("status",1).Filter("article_id",int_aid).Update(orm.Params{"status":0})
		if err != nil {
			return err
		}
fmt.Println("取消赞成功")
		//文章也要减去点赞人数

		_,err = db.QueryTable("gh_article").Filter("id",int_aid).Update(orm.Params{"like_num":orm.ColValue(orm.ColMinus,1)})

		if err != nil {
			return err
		}

		return nil

	}else {
		//没有点赞  就去点赞
		fmt.Println("没有点赞，去点赞")

		//开启事务处理   不会写
		//shiwuErr := db.Begin()

		_,err := db.QueryTable("gh_article").Filter("id",int_aid).Update(orm.Params{"like_num":orm.ColValue(orm.ColAdd,1)})

		if err != nil{
			return err
		}
		fmt.Println("点赞成功")


		like := new(GhLike)
		like.Status = 1
		like.UserId = uid
		like.ArticleId = int_aid

		_,err = db.Insert(like)
		if err != nil {
			return err
		}

		return  nil
	}
}

func CheckLike(uid int64, aid string) bool  {
	int_aid,_ := strconv.ParseInt(aid,10,64)

	cnt,err := db.QueryTable("gh_like").Filter("article_id",int_aid).Filter("user_id",uid).Filter("status",1).Count()
	if err != nil {
		beego.Error(err)
		return false
	}

	if cnt < 1 {
		 return false
	}else {
		return true
	}
}