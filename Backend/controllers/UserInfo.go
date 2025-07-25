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

	if err := global.Db.Where("username = ?", username).First(&user).Error; err != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "用户不存在", nil)
		return
	}
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "参数有误", nil)
		return
	}
	if newUserInfo.NickName == user.NickName || newUserInfo.NickName == "" {
		utils.Respond(c, http.StatusOK, 200, "新昵称不能为空或与旧名称相同！", nil)
	}
	global.Db.Model(&user).Updates(newUserInfo)
	utils.Respond(c, http.StatusOK, 200, "修改成功！", newUserInfo)
}
