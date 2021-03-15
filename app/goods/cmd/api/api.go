package main

import (
	"github.com/spf13/viper"
	"goods/bootstrap"
	"goods/logger"
	"goods/router"
	"net/http"
	"time"
)

func main() {
	// 1. 初始化工作
	bootstrap.Init()
	defer func() {

	}()
	// 2. 启动API服务
	if !viper.IsSet("server") {
		logger.Panic("缺失服务器配置[server]")
	}
	serverViper := viper.Sub("server")
	r := router.InitApiRouter()
	s := &http.Server{
		Addr:              ":" + serverViper.GetString("port"),
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * serverViper.GetDuration("read-timeout"),
		WriteTimeout:      time.Second * serverViper.GetDuration("write-timeout"),
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}
