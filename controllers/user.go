package controllers

import (
	"GinNGorm/logger"
	"GinNGorm/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	user := models.GetUser(id)

	ReturnSuccess(c, 0, "success", user, 1)

}

func (u UserController) GetUserList(c *gin.Context) {

	logger.Write("log message", "user")

	ReturnError(c, 404, "irrelevant information")

}

func (u UserController) AddUser(c *gin.Context) {

	add := &models.User{}
	err := c.BindJSON(&add)
	if err != nil {
		ReturnError(c, 404, gin.H{"Err": "err"})
		return
	}

	id, err := models.AddUser(add.Name, add.Contact, add.Payment, add.History, add.Booking, add.Type)
	if err != nil {
		logger.Write("log message", "user")
		return
	}

	ReturnSuccess(c, 0, "success", id, 1)

}

func (u UserController) UpdateUser(c *gin.Context) {

	user := &models.User{}

	err := c.BindJSON(&user)
	if err != nil {
		ReturnError(c, 404, gin.H{"Err": "err"})
		return
	}

	updateUser, err := models.UpdateUser(user)
	if err != nil {

		logger.Write("log message", "user")

		return
	}

	ReturnSuccess(c, 0, "success", updateUser, 1)

}

func (u UserController) DeleteUser(c *gin.Context) {

	idstr := c.Param("id")
	parseInt, _ := strconv.Atoi(idstr)
	deleteUser, err := models.DeleteUser(parseInt)
	if err != nil {

		ReturnError(c, 404, gin.H{"Err": "err"})
		return
	}

	ReturnSuccess(c, 0, "success", deleteUser, 1)

}
