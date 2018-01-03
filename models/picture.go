package models

import (
	"time"
	"github.com/astaxie/beego"
)

type GhPicture struct {
	Id 			int64		//主键
	Path   	 	string		//本地路径
	Remark 		string		//备注
	Suffix		string		`orm:"null"`	//图片后缀
	CreateTime 	time.Time 	`orm:"auto_now_add;type(datetime)"` //创建时间
}

func PictureInster( pic *GhPicture) (int64, bool) {
	pid,err := db.Insert(pic)
	if err != nil {
		return pid,false
	}
	return pid,true
}

/*
	根据Pid 找对应的图片路径
*/
func GetPictureWithId(pid int64)(path string,err error)  {
	var picture GhPicture
	err = db.QueryTable("gh_picture").Filter("id",pid).One(&picture,"Path")
	if err != nil {
		return "",err
	}else {
		host := beego.AppConfig.String("httphost")
		prot := beego.AppConfig.String("httpport")

		hostPath := "http://" + host + ":" + prot + picture.Path
		return hostPath,nil
	}
}