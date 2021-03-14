package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// CorsMiddleware 跨域控制
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		coreViper := viper.Sub("core")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", coreViper.GetString("origin"))
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", coreViper.GetString("methods"))
			c.Header("Access-Control-Allow-Headers",
				"Content-Type, AccessToken, X-CSRF-Token, Authorization, " + viper.GetString("headerName"))
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
