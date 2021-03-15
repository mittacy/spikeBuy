package router

import (
	"buy/app/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitApiRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	relativePath := "/api/" + viper.GetString("server.version")
	// 初始化控制器
	api := r.Group(relativePath)
	{
		api.POST("/spike/buy", controller.Buy)
		api.POST("/spike/cache", controller.CacheSpike)
	}

	return r
}
