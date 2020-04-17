package db

import (
	"github.com/spf13/viper"
)

var (
	Addrs    = "111.31.105.158:27017"
	User     = "cosmos"
	Passwd   = "cosmos123"
	Database = "cosmos"
)

// 从配置文件初始化
func Init() {
	Addrs = viper.GetString("db.addrs")
	User = viper.GetString("db.user")
	Passwd = viper.GetString("db.passwd")
	Database = viper.GetString("db.database")
}
