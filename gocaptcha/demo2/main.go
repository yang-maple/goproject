package main

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Captcha struct {
	CaptchaId       string `form:"captchaId"`
	CaptchaSolution string `form:"captchaSolution"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", showImage)
	r.GET("/captcha/image.png", showFormHandler)
	r.GET("/captcha/verify", processFormHandler)
	r.Run(":8081")
}

func showImage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func showFormHandler(c *gin.Context) {
	w, h := 240, 80
	captchaId := captcha.New()
	//重写serverHTTP 请求
	_ = serverHTTP(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
	fmt.Println(captchaId)
}

func processFormHandler(c *gin.Context) {
	param := new(Captcha)
	c.Bind(&param)
	fmt.Println(param.CaptchaId, param.CaptchaSolution)
	if !captcha.VerifyString(param.CaptchaId, param.CaptchaSolution) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "succeed",
		})
	}
}

func serverHTTP(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	var captchaId string
	if id != "" {
		captchaId = id
	}
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, captchaId, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, captchaId, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
