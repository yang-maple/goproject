package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	"kubeops/service"
	"net/http"
)

var Daemonset daemonset

type daemonset struct {
}

// 获取list
func (d *daemonset) GetDaemonsetlist(c *gin.Context) {

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
	data, err := service.Daemonset.GetDsList(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Info("获取 Daemonset 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取Daemonset失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// 获取deamonset 详情
func (d *daemonset) GetDaemonsetDetal(c *gin.Context) {
	params := new(struct {
		Daemonsetname string `form:"daemonset_name"`
		Namespace     string `form:"namespace"`
	})

	err := c.Bind(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	Ds, err := service.Daemonset.GetDsDetal(params.Namespace, params.Daemonsetname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("获取Daemonset 失败" + err.Error()),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功",
			"data": Ds,
		})
	}
}

// 删除 实例
func (d *daemonset) DelDaemonset(c *gin.Context) {
	params := new(struct {
		Daemonsetname string `form:"daemonset_name"`
		Namespace     string `form:"namespace"`
	})
	err := c.Bind(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	err = service.Daemonset.DelDs(params.Namespace, params.Daemonsetname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("删除失败" + err.Error()),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
	}

}

// 更新实例
func (d *daemonset) UpdataDaemonset(c *gin.Context) {

	params := new(struct {
		Namespace string            `json:"namespace"`
		Data      *appsv1.DaemonSet `json:"data"`
	})

	err := c.ShouldBindJSON(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	err = service.Daemonset.UpdataDs(params.Namespace, params.Data)
	if err != nil {
		logger.Info("更新失败" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "更新失败",
			"err": errors.New(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})

}
