package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"somnus-gin/global"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化zap
	return zap.New(getZapCore(), options...)
}
