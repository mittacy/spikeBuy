package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}

func Failed(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}



