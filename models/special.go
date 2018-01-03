package models

import "time"

/*
	专题   例如 beego ngrok 等框架专题
*/
type GhSpecial struct {
	Id         int64     //主键
	Status     int64     `orm:"default(1)"` //状态
	UserId     int64     //申请人id
	Tag        string    //标签
	Describe   string    //描述
	Img        string    //专题图标
	HeadImg    string    //专题内部图标
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}
