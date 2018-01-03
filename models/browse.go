package models

import "time"

/*
	浏览记录
*/
type GhBrowse struct {
	Id 			int64			//主键
	Status 		int64			//状态  （清除、隐藏浏览记录 可以用到）
	UserId 		int64			//用户id
	ArticleId 	int64			//文章id
	Tags 		string			//标签组
	CreateTime 	time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}
