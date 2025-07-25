package router

import (
	"Backend/controllers"
	"Backend/middle"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	cate := r.Group("/api/cate")
	{
		cate.GET("/getAllCate", controllers.GetAllCategory)
	}
	user := r.Group("/api/user")
	user.Use(middle.AuthTokenMiddle)
	{
		user.PUT("/update", controllers.UpdateUserInfo)
	}
	return r
}
