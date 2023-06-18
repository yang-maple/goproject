package main

import (
	"github.com/gin-gonic/gin"
	"prometheus/ginx/router"
	_ "prometheus/gmysql"
)

func main() {

	r := gin.Default()
	router.DefRouter(r)
	r.Run()
}
