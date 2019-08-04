package main

import (
	"github.com/qiankaihua/ginDemo/Boot/Config"
	"github.com/qiankaihua/ginDemo/Boot/Http"
	"github.com/qiankaihua/ginDemo/Boot/Log"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
	"github.com/qiankaihua/ginDemo/Route"
)

func _init() {
	Config.InitConfig()
	Log.InitLog()
	Orm.InitOrm()
	// use Orm.GetDB to get orm instance
	//db := Orm.GetDB()
	//db.HasTable("test")
	Log.PushNotice("a", "123")
	Log.PushNotice("b", "abc")
	Log.PrintNotice()
	Http.InitHttp()
	Route.AddApiRoute()
}
func _end() {
	Orm.EndOrm()
}

func main() {
	_init()
	Http.Run()
	_end()
}