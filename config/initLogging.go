package config

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var Log = logrus.New()
var LogFile *os.File

//初始化日志
func init() {
	logName := path.Join(LogFilePath, LogFileName)

	//写入文件
	var err error
	LogFile, err = os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//检查是否发生错误
	if err != nil {
		fmt.Println("err", err)
		return
	}

	//设置日志输出至文件和控制台
	output := io.MultiWriter(os.Stdout, LogFile)
	Log.SetOutput(output)

	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Log.SetFormatter(&logrus.TextFormatter{DisableColors: false, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})

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
		fmt.Println("logger initialization failed!")
		return
	}

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	//建立钩子函数
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{DisableColors: false, FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})

	//添加钩子函数
	Log.AddHook(hook)

	//成功启动logger
	Log.Info("Logger initialized successfully!")
}
