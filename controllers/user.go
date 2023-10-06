package controllers

import (
	"GinNGorm/logger"
	"GinNGorm/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (u UserController) GerUserInfo(c *gin.Context) {
	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	user := models.GetUserTest(id)

	ReturnSuccess(c, 0, "success", user, 1)

}

func (u UserController) GetUserList(c *gin.Context) {

	logger.Write("log message", "user")

	ReturnError(c, 404, "irelevant information")

}
