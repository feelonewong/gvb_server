package router

import "gvb_server/api"

func  (router *RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingsApi.SettingsInfoView)
}