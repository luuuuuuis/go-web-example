/*
@File    :   log.go
@Time    :   2024/04/09 22:03:20
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package utils

// LogInit 完成Zero 日志的初始化

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger zerolog.Logger
)

// LogInit 完成Zero 日志的初始化
func LogInit() {

	// 设置日志等级
	switch Config.GetString("Log.level") {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}

	// 日志切割
	logRotate := &lumberjack.Logger{
		Filename:   strings.Join([]string{Config.GetString("Log.logdir"), Config.GetString("Log.logfile")}, "/"), // 文件位置
		MaxSize:    Config.GetInt("Log.maxsize"),                                                                        // megabytes，M 为单位，达到这个设置数后就进行日志切割
		MaxBackups: Config.GetInt("Log.maxbackups"),                                                                     // 保留旧文件最大份数
		MaxAge:     Config.GetInt("Log.maxage"),                                                                         //days ， 旧文件最大保存天数
		Compress:   Config.GetBool("Log.compress"),                                                                      // disabled by default，是否压缩日志归档，默认不压缩
	}
	// 调整日志时间格式
	zerolog.TimeFieldFormat = time.StampMilli
	// 修改错误日志字段由默认的error修改为err
	zerolog.ErrorFieldName = "err"
	// 打印错误堆栈信息
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// var output = logRotate

	Logger = zerolog.New(logRotate).With().Timestamp().Caller().Logger()

	// 判断环境，开发环境则同时在日志文件和终端打印，生产环境则只在日志文件中打印
	if Config.GetString("Log.pattern") == "development" {
		// 控制台输出的输出器
		// 美化控制台输出，会消耗一定性能，生产环境不建议使用
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		output := zerolog.MultiLevelWriter(consoleWriter, logRotate)
		// log.Logger = log.Output(multi)
		Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	}	
}
