package main

import (
	"somnus-gin/bootstrap"
	"somnus-gin/global"
)

func main() {

	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitValidator()

	// 初始化redis
	global.App.Redis = bootstrap.InitRedis()

	// 启动服务器
	bootstrap.RunServer()
}
