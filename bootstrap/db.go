package bootstrap

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"somnus-gin/global"
)

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	}
}

func getGormLogWirter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFileName,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxAge:     global.App.Config.Log.MaxAge,
			MaxBackups: global.App.Config.Log.MaxBackups,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 writer
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}
