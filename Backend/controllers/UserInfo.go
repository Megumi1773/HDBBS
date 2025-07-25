package controllers

import (
	"Backend/global"
	"Backend/models"
	"Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUserInfo(c *gin.Context) {
	var user models.User
	var newUserInfo models.UpdateUser
	username, _ := c.Get("username")
	global.Db.Where("username = ?", username).First(&user)
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "参数有误", nil)
		return
	}
	if user.NickName == newUserInfo.NickName {
		utils.Respond(c, http.StatusBadRequest, 400, "新昵称不能与旧名称相同！", nil)
		return
	}
	global.Db.Model(&user).Updates(newUserInfo)
	utils.Respond(c, http.StatusOK, 200, "修改成功！", newUserInfo)
}
