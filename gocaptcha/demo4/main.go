package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/wonderivan/logger"
	"image/color"
	"net/http"
)

type configCaptcha struct {
	Id          string `json:"id"`
	VerifyValue string `json:"VerifyValue"`
}

var store = base64Captcha.DefaultMemStore

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	r.GET("/", showImage)
	r.POST("/api/getCaptcha", showFormHandler)
	r.POST("/api/verifyCaptcha", verifyFormHandler)
	r.Run(":8081")
}

func showImage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func showFormHandler(c *gin.Context) {
	params := new(configCaptcha)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		logger.Info("绑定失败")
	}
	driverConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 1,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	captcha := base64Captcha.NewCaptcha(driverConfig.ConvertFonts(), store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      1,
		"data":      b64s,
		"captchaId": id,
		"msg":       "success",
	})
}

func verifyFormHandler(c *gin.Context) {
	params := new(configCaptcha)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		logger.Info("绑定失败")
	}
	fmt.Println(params.Id, params.VerifyValue)
	if store.Verify(params.Id, params.VerifyValue, true) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  "failed",
		})
	}
}
