package routers

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Test)
	r.POST("/admin/login", controllers.Login)
	r.Run(":8080")
	return r
}
