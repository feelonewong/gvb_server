package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化log
	global.Log = core.InitLogger()
	global.Log.Info("info 提示")
	global.Log.Warnln("嘻嘻嘻哈哈哈")
	global.Log.Error("鱻鱼鱼笑脸瞎")
	// 初始化数据库
	global.DB = core.InitGorm()
}
