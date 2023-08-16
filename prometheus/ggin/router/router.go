package router

import (
	"github.com/gin-gonic/gin"
	"prometheus/ggin/controller"
)

func DefRouter(r *gin.Engine) {
	alert := r.Group("/alert")
	{
		alert.POST("/", controller.Alertrequest)
	}
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
}
