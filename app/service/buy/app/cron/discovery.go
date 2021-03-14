package cron

import (
	"buy/logger"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var iclient naming_client.INamingClient

// 我通过example的源码 创建一个真正的注册中心
func InitNacos() {
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	serverConfig := viper.Sub("discovery.server")
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      serverConfig.GetString("ip"),
			ContextPath: "/nacos",
			Port:        serverConfig.GetUint64("port"),
			Scheme:      "http",
		},
	}

	var err error
	iclient, err = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfigs": clientConfig,
	})
	if err != nil {
		logger.PanicDetail("初始化服务注册失败", err)
	}

	instanceViper := viper.Sub("discovery.instance")
	success, err := iclient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          instanceViper.GetString("ip"),
		Port:        viper.GetUint64("server.port"),
		ServiceName: instanceViper.GetString("name"),
		Weight:      instanceViper.GetFloat64("weight"),
		Enable:      true,
		Healthy:     instanceViper.GetBool("healthy"),
		Ephemeral:   instanceViper.GetBool("ephemeral"),
		ClusterName: instanceViper.GetString("cluster-name"),
		GroupName:   instanceViper.GetString("group-name"),
	})
	if err != nil || !success {
		logger.PanicDetail("服务注册失败", err)
	}
	logger.Info("服务注册成功")
}

func DeregisterInstance() {
	logger.Info("服务注销")
	instanceViper := viper.Sub("discovery.instance")
	success, err := iclient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          instanceViper.GetString("ip"),
		Port:        viper.GetUint64("server.port"),
		ServiceName: instanceViper.GetString("name"),
		Ephemeral:   instanceViper.GetBool("ephemeral"),
		Cluster:     instanceViper.GetString("cluster-name"),
		GroupName:   instanceViper.GetString("group-name"),
	})
	if err != nil || !success {
		logger.PanicDetail("服务注销失败", err)
	}
}
