package main

import (
	_ "GoHold/routers"

	"GoHold/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	orm.RunSyncdb("default", false, true)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()

}

func init() {
	//models.RegisterDBS()
	models.Init()

}

type People interface {
	Show()
}
type student struct {
	Name string
	Age  int
}

//
//func (stu *Student) Speak(think string) (talk string){
//	if think == "bitch" {
//		talk = "you are a good boy"
//	}else {
//		talk = "hi"
//	}
//	return
//}

//func (stu *Student)Show()  {
//
//}
//
//func live() People{
//	var stu *Student
//	return stu
//}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
