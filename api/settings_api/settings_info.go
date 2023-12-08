package settings_api

import "github.com/gin-gonic/gin"

func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "响应数据啦啦啦",
	})
}
