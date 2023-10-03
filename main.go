package main

import (
	"GinNGorm/router"
)

func main() {

	r := router.Router()

	//r := gin.Default()
	//
	//r.GET("/hello", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "hello world",
	//	})
	//})
	//
	//
	//user:=r.Group("/user")
	//{
	//
	//	user.POST("/update", func(context *gin.Context) {
	//
	//		context.JSON(http.StatusOK, gin.H{
	//			"message":"update user",
	//		})
	//
	//	})
	//
	//	user.PUT("/add", func(context *gin.Context) {
	//
	//		context.JSON(http.StatusOK, gin.H{
	//			"message":"add user",
	//		})
	//
	//	})
	//	user.DELETE("/delete", func(context *gin.Context) {
	//
	//		context.JSON(http.StatusOK, gin.H{
	//			"message":"delete user",
	//		})
	//
	//	})
	//
	//
	//}

	r.Run(":9999")

}
