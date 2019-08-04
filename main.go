package main

import (
	"fmt"
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
	orm := Orm.GetDB()
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