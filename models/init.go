package models

import (
	"net/url"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)
var db orm.Ormer

func Init() {

	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	// fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(GhArticle),new(GhPicture),new(GhZone),new(GhUser),new(GhSsr),new(GhSpecial),new(GhCollect),new(GhComment),new(GhLike),new(GhBrowse))

	db = orm.NewOrm()
	orm.Debug = true

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true

	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
