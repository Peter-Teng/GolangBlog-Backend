package config

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Log = logrus.New()

//初始化日志
func init() {
	logName := path.Join(LogFilePath, LogFileName)

	//设置日志级别
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		logrus.Error("Log level parsed failed! errMsg = ", err)
		return
	}
	Log.SetLevel(level)

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Error("logger initialization failed! errMsg = ", err)
		return
	}

	//设置log自身Formatter
	Log.SetFormatter(&logrus.TextFormatter{DisableColors: false, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})

	//设置日志输出至文件和控制台
	Log.SetOutput(os.Stdout)

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	//建立钩子函数、设置日志格式
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{DisableColors: false, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})

	//添加钩子函数
	Log.AddHook(hook)

	//成功启动logger
	Log.Info("Logger initialized successfully!")
}
