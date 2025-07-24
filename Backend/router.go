package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "login success",
			})
		})
		auth.POST("/region", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "region success",
			})
		})
	}
	category := r.Group("/api/category")
	{
		category.GET("/list", GetAllCategory)
	}
	return r
}
