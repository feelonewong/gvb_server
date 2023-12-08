package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/router"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化log
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()
	// 初始化路由
	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行: %s", addr)
	r.Run(addr)
}
