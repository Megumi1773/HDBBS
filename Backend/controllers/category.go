package controllers

import (
	"Backend/global"
	"Backend/models"
	"Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllCategory(c *gin.Context) {
	var cate []models.Category
	if res := global.Db.Find(&cate); res.Error != nil {
		utils.Respond(c, http.StatusBadRequest, 400, "获取失败", nil)
		return
	}
	utils.Respond(c, http.StatusOK, 200, "获取成功", cate)

}
