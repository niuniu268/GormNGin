package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

		user.POST("/update", func(context *gin.Context) {

			context.JSON(http.StatusOK, gin.H{
				"message": "update user",
			})

		})

		user.PUT("/add", func(context *gin.Context) {

			context.JSON(http.StatusOK, gin.H{
				"message": "add user",
			})

		})
		user.DELETE("/delete", func(context *gin.Context) {

			context.JSON(http.StatusOK, gin.H{
				"message": "delete user",
			})

		})

	}

	return r
}
