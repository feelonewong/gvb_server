package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		//
		global.Log.Warnln("未配置MySQL，取消gorm链接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境打印所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql 连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大可容纳数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间
	return db
}
