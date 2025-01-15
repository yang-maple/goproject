package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prometheus/gmysql"
	"prometheus/notice"
	"prometheus/pstruct"
	"time"

	"github.com/gin-gonic/gin"
)

//type T struct {
//	Alerts []struct {
//		Annotations struct {
//			Description string `json:"description"`
//			Summary     string `json:"summary"`
//		} `json:"annotations"`
//		EndsAt      time.Time `json:"endsAt"`
//		Fingerprint string    `json:"fingerprint"`
//		Labels      struct {
//			Alertname string `json:"alertname"`
//			Instance  string `json:"instance"`
//			Ip        string `json:"ip"`
//			Job       string `json:"job"`
//			Monitor   string `json:"monitor"`
//			Region    string `json:"region"`
//			Replica   string `json:"replica"`
//			Severity  string `json:"severity"`
//			Port      string `json:"port,omitempty"`
//		} `json:"labels"`
//		StartsAt time.Time `json:"startsAt"`
//		Status   string    `json:"status"`
//	} `json:"alerts"`
//	CommonAnnotations struct {
//		Description string `json:"description"`
//		Summary     string `json:"summary"`
//	} `json:"commonAnnotations"`
//	CommonLabels struct {
//		Alertname string `json:"alertname"`
//		Job       string `json:"job"`
//		Monitor   string `json:"monitor"`
//		Region    string `json:"region"`
//		Replica   string `json:"replica"`
//		Severity  string `json:"severity"`
//	} `json:"commonLabels"`
//	ExternalURL string `json:"externalURL"`
//	GroupKey    string `json:"groupKey"`
//	GroupLabels struct {
//		Alertname string `json:"alertname"`
//	} `json:"groupLabels"`
//	Receiver        string `json:"receiver"`
//	Status          string `json:"status"`
//	TruncatedAlerts int    `json:"truncatedAlerts"`
//	Version         string `json:"version"`
//}

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

func Alertrequest(c *gin.Context) {
	var data pstruct.AlertMessage
	if err := c.ShouldBindJSON(&data); err != nil {
		// 如果绑定失败，则抛出异常
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err := gmysql.DB.AutoMigrate(&gmysql.AlertMessageSql{})
	if err != nil {
		fmt.Println("创建数据表出错", err)
	}
	for _, v := range data.Alerts {
		alerts := gmysql.AlertMessageSql{
			CreatedAt: time.Now(),
			Alerts:    v,
		}
		result := gmysql.DB.Create(&alerts)
		fmt.Println(result)
		datajson, _ := json.Marshal(v)
		fmt.Println(string(datajson))
		notice.WeRebot(&v)
<<<<<<< HEAD
		//	notice.Emails(v, notice.FormatEmailBody("view/notice.html", alerts))
=======
		notice.Emails(v, notice.FormatEmailBody("view/notice.html", alerts))
>>>>>>> 4622a0a (2025-1-15)

	}

	c.JSON(http.StatusOK, "ok")
}
