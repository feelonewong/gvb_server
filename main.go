package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	global.DB = core.InitGorm()
}
