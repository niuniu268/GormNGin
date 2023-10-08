package router

import (
	"GinNGorm/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	user := r.Group("/user")
	{

		//user.POST("/update", func(context *gin.Context) {
		//
		//	context.JSON(http.StatusOK, gin.H{
		//		"message": "update user",
		//	})
		//
		//})
		//
		//user.PUT("/add", func(context *gin.Context) {
		//
		//	context.JSON(http.StatusOK, gin.H{
		//		"message": "add user",
		//	})
		//
		//})
		//user.DELETE("/delete", func(context *gin.Context) {
		//
		//	context.JSON(http.StatusOK, gin.H{
		//		"message": "delete user",
		//	})
		//
		//})
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.PUT("/update", controllers.UserController{}.UpdateUser)
		user.DELETE("/info/:id", controllers.UserController{}.DeleteUser)

	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetOderList)
	}

	//r.Run(":9999")

	return r
}
