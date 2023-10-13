package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubeops/service"
	"net/http"
)

type persistenvolume struct{}

var Persistentvolume persistenvolume

// GetPersistentVolumeList 获取PersistentVolume清单
func (p *persistenvolume) GetPersistentVolumeList(c *gin.Context) {
	data, err := service.Persistenvolume.GetPvList()
	if err != nil {
		logger.Info("获取 PersistentVolume 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取PersistentVolume失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// GetPersistentVolumeDetail 获取 PersistentVolume 详情
func (p *persistenvolume) GetPersistentVolumeDetail(c *gin.Context) {
	params := new(struct {
		PersistentVolumeName string `form:"persistent_volume_name"`
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
	data, err := service.Persistenvolume.GetPvDetail(params.PersistentVolumeName)
	if err != nil {
		logger.Info("获取 PersistentVolume 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取PersistentVolume失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取数据成功",
		"data": data,
	})
}

// DelPersistentVolume 删除 PersistentVolume 资源
func (p *persistenvolume) DelPersistentVolume(c *gin.Context) {
	params := new(struct {
		PersistentVolumeName string `form:"persistent_volume_name"`
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
	err = service.Persistenvolume.DelPv(params.PersistentVolumeName)
	if err != nil {
		logger.Info("删除 PersistentVolume 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "删除 PersistentVolume 失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// CreatePersistentVolume 创建 PersistentVolume 资源
func (p *persistenvolume) CreatePersistentVolume(c *gin.Context) {
	data := new(service.CreatePVConfig)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logger.Info("绑定参数失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "绑定参数失败",
			"data": nil,
		})
		return
	}
	err = service.Persistenvolume.CreatePv(data)
	if err != nil {
		logger.Info("创建 PersistentVolume 失败" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "创建PersistentVolume失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}
