package initSettings

import (
	"VueGin/global"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//設定Init，以後只需要呼叫zap.L()
//returns the global Logger, which can be reconfigured with ReplaceGlobals. It's safe for concurrent use.
func InitLogger() (*zap.Logger, error) {
	level := global.Global_Config.Log.Level
	filename := global.Global_Config.Log.FileName
	maxsize := global.Global_Config.Log.MaxSize
	maxbackup := global.Global_Config.Log.MaxBackups
	maxage := global.Global_Config.Log.MaxAge

	writeSyncer := getLogWriter(filename, maxsize, maxbackup, maxage)
	encoder := getEncorder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(level))
	if err != nil {
		fmt.Println("Counldn't initialize zap logger")
		return nil, nil
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	//將新建立的logger參數，加入全域變數中
	// global.Global_logger = zap.New(core, zap.AddCaller())
	// global.Global_Config	.
	return zap.New(core, zap.AddCaller()), nil
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	tmp := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(tmp)
}

// func getEncorder() zapcore.EncoderConfig {
func getEncorder() zapcore.Encoder {
	//建立生產環境參數
	// cfg := zap.NewProductionEncoderConfig()
	// encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderConfig.TimeKey = "time"
	// encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// return zapcore.NewJSONEncoder(encoderConfig)

	//少一個都不行，會報錯nil
	cfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return zapcore.NewJSONEncoder(cfg)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("gin_server - " + "2006/01/02 - 15:04:05.000"))
}
