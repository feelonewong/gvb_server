package router

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()

	// 路由分组
	apiGroup := r.Group("/api")
	routerGroupApp := RouterGroup{apiGroup}
	// 系统配置
	routerGroupApp.SettingRouter()

	return r
}
