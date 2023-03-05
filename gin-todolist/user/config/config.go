package config

import (
	"os"

	"github.com/spf13/viper"
)

// 用来获取config
func InitConfig() {
	workDir, _ := os.Getwd()                 // 获取当前工作目录, 注: 终端当前所在的目录
	viper.AddConfigPath(workDir + "/config") // 配置文件的文件路径
	viper.SetConfigName("config")            // 配置文件的文件名
	viper.SetConfigType("yaml")              // 配置文件的文件后缀
	if err := viper.ReadInConfig(); err != nil {
		return
	}
}
