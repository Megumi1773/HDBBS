package controllers

import (
	"Backend/global"
	"Backend/models"
	"Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "参数有误", nil)
		return
	}
	if res := global.Db.Where("username = ?", loginUser.Username).First(&user); res.RowsAffected == 0 {
		utils.Respond(c, http.StatusUnauthorized, 401, "用户不存在！", nil)
		fmt.Println(user)
		return
	}
	if user.Password != loginUser.Password {
		utils.Respond(c, http.StatusBadRequest, 400, "密码错误！", nil)
		return
	}
	token, err := utils.CreateToken(user.Username)
	if err != nil {
		log.Fatalf("token create err:%v", err)
	}
	newUser := gin.H{
		"Username": user.Username,
		"token":    token,
	}
	utils.Respond(c, http.StatusOK, http.StatusOK, "登录成功！", newUser)
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "参数有误", nil)
		return
	}
	if res := global.Db.Create(&user); res.Error != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "注册失败", nil)
		return
	}
	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}
	utils.Respond(c, http.StatusOK, http.StatusOK, "注册成功！", newUser)
}
