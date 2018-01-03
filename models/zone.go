package models
//个人空间
type GhZone struct {
	Id int64 //
	Status int64 `orm:"default(1)"` // 1 启用   0 禁用
	Remark string	`orm:"default(懒~)"`
}
