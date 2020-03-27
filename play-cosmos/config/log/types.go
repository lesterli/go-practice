package log

import (
	"github.com/spf13/viper"
)

var (
	LogLevel          = "debug"
	FileName          = "logs/sync_cosmos.log"
	MaxSize           = 20
	MaxAge            = 7
	MaxBackups        = 3
	Compress          = true
	EnableAtomicLevel = true
)

// Init 从配置文件初始化
func Init() {
	LogLevel = viper.GetString("log.level")
	FileName = viper.GetString("log.file")
	MaxSize = viper.GetInt("log.rotate_size")
	MaxAge = viper.GetInt("log.rotate_date")
	MaxBackups = viper.GetInt("log.backup_count")
	EnableAtomicLevel = viper.GetBool("log.enable_atomic_level")
}
