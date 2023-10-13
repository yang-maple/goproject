package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubeops/service"
	"net/http"
)

type namespace struct{}

var Namespace namespace

func (n *namespace) GetNslist(c *gin.Context) {
	data, err := service.Namespace.GetNslist()
	if err != nil {
		logger.Info("获取 namespaces 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取namespaces失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

func (n *namespace) GetNsDetail(c *gin.Context) {
	params := new(struct {
		NamespaceName string `form:"namespace_name"`
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
	data, err := service.Namespace.GetNsDetal(params.NamespaceName)
	if err != nil {
		logger.Info("获取 namespaces 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取namespaces失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

func (n *namespace) DelNs(c *gin.Context) {
	params := new(struct {
		NamespaceName string `form:"namespace_name"`
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
	err = service.Namespace.DelNs(params.NamespaceName)
	if err != nil {
		logger.Info("获取 namespaces 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "删除namespaces失败",
			"data": errors.New("删除失败" + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func (n *namespace) CreateNs(c *gin.Context) {
	params := new(struct {
		NamespaceName string `form:"namespace_name"`
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
	err = service.Namespace.CreateNs(params.NamespaceName)
	if err != nil {
		logger.Info("创建 namespaces 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "创建namespaces失败",
			"data": errors.New("创建失败" + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}
