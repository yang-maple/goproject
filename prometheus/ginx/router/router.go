package router

import (
	"github.com/gin-gonic/gin"
	"prometheus/ginx/controller"
)

func DefRouter(r *gin.Engine) {

	alert := r.Group("/alert")
	//alert.Use(func(c *gin.Context) {
	//	body, err := io.ReadAll(c.Request.Body)
	//	if err != nil {
	//		c.AbortWithStatus(http.StatusInternalServerError)
	//		return
	//	}
	//	c.Set("json", body)
	//	fmt.Println("body :", string(body))
	//
	//})
	{
		alert.POST("/", controller.Alertrequest)
	}
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
}
