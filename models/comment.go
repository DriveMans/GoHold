package models

import (
	"time"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
)

/*
	回复
*/
type GhComment struct {
	Id         	int64     	//主键
	Status     	int64     	`orm:"default(1)"`
	ArticleId  	int64     	//文章id
	UserId     	int64     	//用户id
	Content    	string    	//回复内容
	BeCommentId	int64	 	`orm:"null"`//被回复的评论id 如果为空就是直接评论  否则为间接评论
	BeUserId	int64		`orm:"null"`//被回复人的id 可以在间接评论层回复他人
	Tree       	int64     	`orm:"default(1)"`                  //权重 1 > 2 > 3
	CreateTime 	time.Time 	`orm:"auto_now_add;type(datetime)"` //创建时间
}

/*
	回复文章
*/
func ReplyArticle(comment *GhComment)error  {
	_,err := db.Insert(comment)
	return err
}


/*
	获取所有评论
*/
func GetAllComment(article string)([]*GhComment,error)  {
	int_aid,_ := strconv.ParseInt(article,10,64)
	var comments []*GhComment

	num,err := db.QueryTable("gh_comment").Filter("article_id",int_aid).All(&comments)
	if err != nil  {
		beego.Error(err)
		return nil,err
	}

	fmt.Printf("查询到%d条数据\n",num)
	if num < 1 {
		return nil,fmt.Errorf("暂无评论内容")

	}else {
		return comments,nil

	}
}