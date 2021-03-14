package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件, %s\n", err))
		} else {
			panic(fmt.Errorf("err: %s\n", err))
		}
	}
	fmt.Println("配置文件初始化成功")
}
