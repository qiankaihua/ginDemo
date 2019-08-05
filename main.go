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
	Log.InitTimer()
	Orm.InitOrm()

	// use Orm.GetDB to get orm instance
	//db := Orm.GetDB()
	//db.HasTable("test")

	// append notice, [key=value]
	//Log.PushNotice("a", "123")
	//Log.PushNotice("b", "abc")

	// SetTimer 设置定时器, 设置为现在开始, 再设置会覆盖之前的时间
	//Log.SetTimer("test")

	Http.InitHttp()
	Route.AddStaticRoute()
	Route.AddApiRoute()

	// GetTimer 获取定时器时间
	//Log.PushNotice("test_timer", Log.GetTimer("test"))

	// 停止计时器, 重复停止按第一次计数
	//Log.StopTimer("test")
	//Log.StopAll()

	// 打印 PushNotice 的日志
	//Log.PrintNotice()

	// 设置 Error 级别的 log
	//Log.Error("Error Log")
}
func _end() {
	Orm.EndOrm()
}

func main() {
	_init()
	Http.Run()
	_end()
}