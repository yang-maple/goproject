package controller

import (
	"encoding/json"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"kubeops/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Pod pod

type pod struct {
}

// 获取pod list
func (p *pod) GetPods(c *gin.Context) {
	//定义匿名结构体 处理入参的请求
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	err := c.Bind(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "绑定参数失败",
			"data": nil,
		})
		return
	}
	data, err := service.Pod.GetPods(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Info("获取podlist失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取podlist失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})

}

// 获取pod 详情
func (p *pod) GetPodsDetail(c *gin.Context) {
	params := new(struct {
		Namespace string `form:"namespace"`
		PodName   string `form:"pod_name"`
	})
	c.Bind(&params)
	fmt.Println(params)
	data, err := service.Pod.GetPodDetail(params.PodName, params.Namespace)
	if err != nil {
		logger.Info("获取pod 详情失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取成功",
		"data": data,
	})
}

// 删除 pod
func (p *pod) DeletePod(c *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		PodName   string `json:"pod_name"`
	})
	c.ShouldBind(&params)
	fmt.Println("参数为：", params)
	err := service.Pod.DelPod(params.PodName, params.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
	}
}

// 更新 pod
func (p *pod) UpdataPod(c *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	c.ShouldBindJSON(&params)
	fmt.Println("绑定参数", params.Content)
	//反序列化
	var newpod = &corev1.Pod{}
	err := json.Unmarshal([]byte(params.Content), newpod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		logger.Info("err" + err.Error())
		return
	}
	fmt.Println(newpod)
	err = service.Pod.UpdataPod(newpod, params.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "更新成功",
		})
	}
}

// 获取容器列表
func (p *pod) GetContainerList(c *gin.Context) {
	params := new(struct {
		Namespace string `json:"namespace"`
		PodName   string `json:"pod_name"`
	})

	c.ShouldBind(&params)
	containers, err := service.Pod.GetContainer(params.PodName, params.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取容器信息失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取容器信息成功",
			"data": containers,
		})
	}
}

// 获取容器日志
func (p *pod) GetContainerLog(c *gin.Context) {

	params := new(struct {
		Namespace     string `json:"namespace"`
		PodName       string `json:"pod_name"`
		Containername string `json:"containername"`
	})
	c.ShouldBind(&params)
	log, err := service.Pod.GetContainerLog(params.PodName, params.Containername, params.Namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取日志失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取日志成功",
			"data": log,
		})
	}

}

// 获取每个namespace下的pod 数量
func (p *pod) GetPodnumer(c *gin.Context) {
	total, err := service.Pod.Countpod()
	fmt.Println("total:", total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取数据成功",
			"data": total,
		})
	}
}
