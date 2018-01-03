package models

import (
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)

/*
	文章列表
*/
type GhArticle struct {
	Id         int64     //主键
	Status     int64     `orm:"default(1)"` //状态
	UserId     int64     //用户id
	ViewNum    int64     `orm:"default(0)"` //阅读次数
	LikeNum    int64     `orm:"default(0)"` //点赞数量
	Title      string    //文章标题
	Tags       string    //标签组(beego   ngrok ...)
	Url        string    //文章路径
	Img        string    //文章预览图
	IsHot      bool      `orm:"default(false)"`              //是否热门
	IsEssence  bool      `orm:"default(false)"`              //是否精华
	IsBanner   bool      `orm:"default(false)"`              //是否置顶到轮播展示
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}

/*
	插入文章
*/
func InsertArticle(article *GhArticle) error {
	_, err := db.Insert(article)
	return err
}

/*
	获取首页文章
*/
func GetIndexDataWithIndexPage(indexPage string) ([]*GhArticle, error) {
	var articles []*GhArticle
	index, _ := strconv.ParseInt(indexPage,10,64)
	_, err := db.QueryTable("gh_article").OrderBy("-id").Limit(10, index*10).All(&articles)
	return articles, err
}

/*
	浏览文章
*/
func ViewArticle(aid ,token string)  {
	//只需要记录  不用管成功与否
	int_aid,_ := strconv.ParseInt(aid,10,64)
	uid ,err := TokenExchangeUserId(token)

	if err != nil {
		return
	}

	db.QueryTable("gh_article").Filter("id",int_aid).Update(orm.Params{"view_num":orm.ColValue(orm.ColAdd,1)})

	browse := new(GhBrowse)
	browse.Status = 1
	browse.ArticleId = int_aid
	browse.UserId = uid

	db.Insert(browse)
}