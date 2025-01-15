package router

import (
	"github.com/gin-gonic/gin"
	"kube-prometheus-alert/ggin/controller"
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
