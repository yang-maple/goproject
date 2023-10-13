package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubeops/service"
	"net/http"
)

type services struct{}

var Services services

// GetServiceList 获取 services 列表
func (s *services) GetServiceList(c *gin.Context) {
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
	data, err := service.Services.GetSvcList(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Info("获取Services失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取Services失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// GetServiceDetail 获取 services 详情
func (s *services) GetServiceDetail(c *gin.Context) {
	params := new(struct {
		ServiceName string `form:"service_name"`
		Namespace   string `form:"namespace"`
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
	detail, err := service.Services.GetSvcDetail(params.Namespace, params.ServiceName)
	if err != nil {
		logger.Info("获取Services 详情失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取Services 详情失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": detail,
	})
}

// DelServices  删除 services 资源
func (s *services) DelServices(c *gin.Context) {
	params := new(struct {
		ServiceName string `form:"service_name"`
		Namespace   string `form:"namespace"`
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
	err = service.Services.DelSvc(params.Namespace, params.ServiceName)
	if err != nil {
		logger.Info("删除 Services 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "删除 Services 失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除 Services 成功",
	})
}

// CreateService 创建 Services 资源
func (s *services) CreateService(c *gin.Context) {
	createsvc := new(service.CreateService)
	err := c.ShouldBindJSON(&createsvc)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "绑定参数失败",
			"data": nil,
		})
		return
	}
	err = service.Services.CreateSvc(createsvc)
	if err != nil {
		logger.Info("创建 Services 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "创建 Services 失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "创建 Services 成功",
	})
}
