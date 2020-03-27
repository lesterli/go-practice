package log

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger "加糖"
var Logger *zap.SugaredLogger //global logger

// Init 配置初始化
func Init() {
	logLevel := viper.GetString("log.level")

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.file"),
		MaxSize:    viper.GetInt("log.rotate_size"),
		MaxAge:     viper.GetInt("log.rotate_date"),
		MaxBackups: viper.GetInt("log.backup_count"),
		LocalTime:  true,
		Compress:   true,
	})

	zapLogLevel := zap.NewAtomicLevel()
	if err := zapLogLevel.UnmarshalText([]byte(strings.ToLower(logLevel))); err != nil {
		panic(fmt.Errorf("get config log level:%v config error: %v", logLevel, err))
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zapLogLevel,
	)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
	Logger.Info("logger init successful!")
}
