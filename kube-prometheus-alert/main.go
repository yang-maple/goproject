package main

import (
	"github.com/gin-gonic/gin"
	"kube-prometheus-alert/ggin/router"
)

func main() {

	r := gin.Default()
	router.DefRouter(r)
	r.Run()
}
