package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
