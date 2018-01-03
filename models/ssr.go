package models

import (
	"time"
	"fmt"
)

/*
	订阅表
*/
type GhSsr struct {
	Id     		int64  //主键
	Status     	int64     `orm:"default(1)"` //状态 1订阅中   0 解除订阅
	Tag 		string	//订阅标签 归类
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`	//订阅时间
	UpdateTime  time.Time `orm:"null"`	//更新时间
	SubjectId   int64 	//发布者 （被观察者）
	ObserverId	int64	//观察者
}

/*
	订阅
*/
func SubScribe(be_userid,user_id int64) (error) {
	var ssr = new(GhSsr)
	ssr.ObserverId = user_id
	ssr.SubjectId = be_userid

	_,err := db.Insert(ssr)
	return err
}

func SubAllData(user_id int64)([]*GhSsr,error){
	var ssrArr []*GhSsr
	
	num ,err := db.QueryTable("gh_ssr").Filter("observer_id",user_id).All(&ssrArr)
	if err != nil {
		return nil,err
	}

	if num < 1 {
		return nil,fmt.Errorf("暂无数据")
	}

	return ssrArr,nil
}

