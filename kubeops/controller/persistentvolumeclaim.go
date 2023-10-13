package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubeops/service"
	"net/http"
)

type persistentvolumeclaim struct{}

var PersistentVolumeClaim persistentvolumeclaim

// GetPersistentVolumeClaimList 获取 PersistentVolumeClaim 的列表
func (p *persistentvolumeclaim) GetPersistentVolumeClaimList(c *gin.Context) {
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
	data, err := service.Pvclaim.GetPVClaimList(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Info("获取PersistentVolumeClaim失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取PersistentVolumeClaim失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// GetPersistentVolumeClaimDetail 获取 PersistentVolumeClaim 详情
func (p *persistentvolumeclaim) GetPersistentVolumeClaimDetail(c *gin.Context) {
	params := new(struct {
		PersistentVolumeClaimName string `form:"persistent_volume_claim_name"`
		Namespace                 string `form:"namespace"`
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
	detail, err := service.Pvclaim.GetPVClaimDetail(params.Namespace, params.PersistentVolumeClaimName)
	if err != nil {
		logger.Info("获取PersistentVolumeClaim 详情失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取PersistentVolumeClaim 详情失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": detail,
	})
}

// CreatePersistentVolumeClaim 创建PersistentVolumeClaim 资源
func (p *persistentvolumeclaim) CreatePersistentVolumeClaim(c *gin.Context) {
	createpvc := new(service.Createpvclaim)
	err := c.ShouldBindJSON(&createpvc)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "绑定参数失败",
			"data": nil,
		})
		return
	}
	err = service.Pvclaim.CreatePVClaim(createpvc)
	if err != nil {
		logger.Info("创建PersistentVolumeClaim失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "创建PersistentVolumeClaim失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "创建PersistentVolumeClaim成功",
	})
}

// DelPersistentVolumeClaim 删除 PersistentVolumeClaim 资源
func (p *persistentvolumeclaim) DelPersistentVolumeClaim(c *gin.Context) {
	params := new(struct {
		PersistentVolumeClaimName string `form:"persistent_volume_claim_name"`
		Namespace                 string `form:"namespace"`
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
	err = service.Pvclaim.DelPVClaim(params.Namespace, params.PersistentVolumeClaimName)
	if err != nil {
		logger.Info("删除 PersistentVolumeClaim 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "删除 PersistentVolumeClaim 失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除 PersistentVolumeClaim 成功",
	})
}
