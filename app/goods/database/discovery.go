package database

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"goods/logger"
	"strconv"
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
}

func GetAllBuyInstance() []string {
	services, err := iclient.GetService(vo.GetServiceParam{
		ServiceName: "spike.buy",
		Clusters:    []string{"spike-buy"}, // 默认值DEFAULT
		GroupName:   "spike-buy-group",             // 默认值DEFAULT_GROUP
	})
	if err != nil {
		logger.PanicDetail("获取所有抢购服务错误", err)
	}
	if len(services.Hosts) == 0 {
		logger.Error("没有抢购服务")
		return []string{}
	}

	hosts := make([]string, len(services.Hosts))
	for i, v := range services.Hosts {
		hosts[i] = v.Ip + ":" + strconv.Itoa(int(v.Port))
	}
	return hosts
}
