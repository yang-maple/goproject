package service

//列表
// 获取node 详情
import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type persistenvolume struct{}

var Persistenvolume persistenvolume

type pvlist struct {
	Total               int                       `json:"total"`
	Persistenvolumename []corev1.PersistentVolume `json:"persistenvolumename"`
}

type CreatePVConfig struct {
	Name    string            `json:"name"`
	Labels  map[string]string `json:"labels"`
	Storage string            `json:"storage"`
	Type    string            `json:"type"`
	Path    string            `json:"path"`
	Server  string            `json:"server"`
}

// GetPvList PersistentVolume列表
func (p *persistenvolume) GetPvList() (*pvlist, error) {
	pvs, err := K8s.Clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取PersistentVolume 失败" + err.Error())
		return nil, errors.New("获取PersistentVolume 失败" + err.Error())
	}
	item := make([]corev1.PersistentVolume, 0, len(pvs.Items))
	for i := range pvs.Items {
		item = append(item, pvs.Items[i])
	}
	return &pvlist{
		Total:               len(pvs.Items),
		Persistenvolumename: item,
	}, nil
}

// GetPvDetail 获取PersistentVolume 详情
func (p *persistenvolume) GetPvDetail(PvName string) (detail *corev1.PersistentVolume, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().PersistentVolumes().Get(context.TODO(), PvName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取PersistentVolume 详情失败" + err.Error())
		return nil, errors.New("获取PersistentVolume 详情失败" + err.Error())
	}
	return detail, nil
}

// DelPv 删除 PersistentVolume
func (p *persistenvolume) DelPv(PvName string) (err error) {
	err = K8s.Clientset.CoreV1().PersistentVolumes().Delete(context.TODO(), PvName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除 Persistenvolume 详情失败" + err.Error())
		return errors.New("删除 Persistenvolume 详情失败" + err.Error())
	}
	return nil
}

// CreatePv 创建 PersistentVolume
func (p *persistenvolume) CreatePv(data *CreatePVConfig) (err error) {
	var mode corev1.PersistentVolumeMode = corev1.PersistentVolumeFilesystem
	createpv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: corev1.ResourceList{
				corev1.ResourceStorage: resource.MustParse(data.Storage),
			},
			AccessModes:                   []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			ClaimRef:                      nil,
			PersistentVolumeReclaimPolicy: "",
			StorageClassName:              "",
			MountOptions:                  nil,
			VolumeMode:                    &mode,
			NodeAffinity:                  nil,
		},
	}
	switch data.Type {
	case "NFS":
		createpv.Spec.NFS = &corev1.NFSVolumeSource{
			Server: data.Server,
			Path:   data.Path,
		}
		createpv.Spec.StorageClassName = "nfs-client"
	case "HostPATH":
		createpv.Spec.HostPath = &corev1.HostPathVolumeSource{
			Path: data.Path,
			Type: nil,
		}
		createpv.Spec.StorageClassName = "standard"
	}
	_, err = K8s.Clientset.CoreV1().PersistentVolumes().Create(context.TODO(), createpv, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 PersistentVolume 失败" + err.Error())
		return errors.New("创建 PersistentVolume 失败" + err.Error())
	}
	return nil
}
