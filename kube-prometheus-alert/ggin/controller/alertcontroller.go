package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type T struct {
	Alerts []struct {
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		EndsAt      time.Time `json:"endsAt"`
		Fingerprint string    `json:"fingerprint"`
		Labels      struct {
			Alertname string `json:"alertname"`
			Instance  string `json:"instance"`
			Ip        string `json:"ip"`
			Job       string `json:"job"`
			Monitor   string `json:"monitor"`
			Region    string `json:"region"`
			Replica   string `json:"replica"`
			Severity  string `json:"severity"`
			Port      string `json:"port,omitempty"`
		} `json:"labels"`
		StartsAt time.Time `json:"startsAt"`
		Status   string    `json:"status"`
	} `json:"alerts"`
	CommonAnnotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Job       string `json:"job"`
		Monitor   string `json:"monitor"`
		Region    string `json:"region"`
		Replica   string `json:"replica"`
		Severity  string `json:"severity"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	Receiver        string `json:"receiver"`
	Status          string `json:"status"`
	TruncatedAlerts int    `json:"truncatedAlerts"`
	Version         string `json:"version"`
}

//type AlertMessage struct {
//	Alerts []Alerts `json:"alerts"`
//}
//
//type Alerts struct {
//	Annotations struct {
//		Description string `json:"description"`
//		Summary     string `json:"summary"`
//	} `json:"annotations"`
//	Labels struct {
//		Alertname string `json:"alertname"`
//		Instance  string `json:"instance"`
//		Ip        string `json:"ip"`
//		Job       string `json:"job"`
//		Region    string `json:"region"`
//		Severity  string `json:"severity"`
//		Port      string `json:"port,omitempty"`
//	} `json:"labels"`
//	Status string `json:"status"`
//}

type data interface {
}

func Alertrequest(c *gin.Context) {
	var datas T
	err := c.ShouldBindJSON(&datas)
	if err != nil {
		fmt.Println("解析异常")
	}
	fmt.Println(datas)
	c.JSON(http.StatusOK, "ok")
}
