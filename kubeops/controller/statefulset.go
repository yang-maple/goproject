package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	"kubeops/service"
	"net/http"
)

var Statefulset statefulset

type statefulset struct {
}

// 获取list
func (s *statefulset) GetStatefulSetList(c *gin.Context) {

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
	data, err := service.Statefulset.GetStsList(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Info("获取 StatefulSet 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取StatefulSet失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// 获取Statefulset 详情
func (s *statefulset) GetStatefulSetDeTal(c *gin.Context) {
	params := new(struct {
		StatefulSetName string `form:"statefulset_name"`
		Namespace       string `form:"namespace"`
	})

	err := c.Bind(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	Ds, err := service.Statefulset.GetStsDetal(params.Namespace, params.StatefulSetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("获取Statefulset 失败" + err.Error()),
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
func (s *statefulset) DelStatefulSet(c *gin.Context) {
	params := new(struct {
		StatefulSetName string `form:"statefulset_name"`
		Namespace       string `form:"namespace"`
	})
	err := c.Bind(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	err = service.Statefulset.DelSts(params.Namespace, params.StatefulSetName)
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
func (s *statefulset) UpDataStatefulSet(c *gin.Context) {

	params := new(struct {
		Namespace string              `json:"namespace"`
		Data      *appsv1.StatefulSet `json:"data"`
	})

	err := c.ShouldBindJSON(&params)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": errors.New("绑定参数失败" + err.Error()),
		})
		return
	}
	err = service.Statefulset.UpdataSts(params.Namespace, params.Data)
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
