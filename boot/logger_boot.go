package boot

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type LoggerBootloader struct {
}

func NewLoggerBootloader() *LoggerBootloader {
	return &LoggerBootloader{}
}

func (sb LoggerBootloader) Boot() error {
	//fileName := application.AppSetting.LogSavePath + "/" + application.AppSetting.LogFileName + application.AppSetting.LogFileExt
	fileName := "./test.log"
	application.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: false,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
