package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type pvclaim struct{}
type Pvclaimsresp struct {
	Total int                            `json:"total"`
	Item  []corev1.PersistentVolumeClaim `json:"item"`
}

type Createpvclaim struct {
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	Labels           map[string]string `json:"labels"`
	AccessModes      string            `json:"access_modes"`
	Storage          string            `json:"storage"`
	StorageClassName string            `json:"storage_classname" default:"nfs-client"`
}

var Pvclaim pvclaim

func (p pvclaim) tocells(pvclaims []corev1.PersistentVolumeClaim) []DataCell {
	cells := make([]DataCell, len(pvclaims))
	for i := range pvclaims {
		cells[i] = pvcCell(pvclaims[i])
	}
	return cells
}

func (p pvclaim) fromcells(cells []DataCell) []corev1.PersistentVolumeClaim {
	pvclaims := make([]corev1.PersistentVolumeClaim, len(cells))
	for i := range cells {
		pvclaims[i] = corev1.PersistentVolumeClaim(cells[i].(pvcCell))
	}
	return pvclaims
}

// 列表
func (p *pvclaim) GetPVClaimList(Pvclaimname, Namespace string, Limit, Page int) (DP *Pvclaimsresp, err error) {
	//获取deployment 的所有清单列表
	pvclaimlist, err := K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 PersistentVolumeClaims 失败" + err.Error())
	}
	//组装数据
	selectdata := &dataselector{
		GenericDataList: p.tocells(pvclaimlist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Pvclaimname},
			Paginate: &PaginateQuery{
				limit: Limit,
				page:  Page,
			},
		},
	}
	//先过滤 后排序
	filtered := selectdata.Filter()
	total := len(filtered.GenericDataList)
	//分页
	datapage := filtered.Sort().Pagination()
	pvclaims := p.fromcells(datapage.GenericDataList)
	for _, v := range pvclaims {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &Pvclaimsresp{
		Total: total,
		Item:  pvclaims,
	}, err
}

// 详情
func (p *pvclaim) GetPVClaimDetail(Namespace, Pvclaimname string) (detail *corev1.PersistentVolumeClaim, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).Get(context.TODO(), Pvclaimname, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取 PersistentVolumeClaims 详情失败" + err.Error())
		return nil, errors.New("获取PersistentVolumeClaims 详情失败" + err.Error())
	}
	return detail, nil
}

// 创建
func (p *pvclaim) CreatePVClaim(data *Createpvclaim) (err error) {

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Labels,
			},
			StorageClassName: &data.StorageClassName,
			Resources: corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse(data.Storage),
				},
			},
		},
		Status: corev1.PersistentVolumeClaimStatus{},
	}

	_, err = K8s.Clientset.CoreV1().PersistentVolumeClaims(data.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 secret 失败" + err.Error())
		return errors.New("创建 secret 失败" + err.Error())
	}
	return nil

}

// 删除
func (p *pvclaim) DelPVClaim(Namespace, Pvclaimname string) (err error) {
	err = K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).Delete(context.TODO(), Pvclaimname, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除Pvclaimname失败" + err.Error())
		return errors.New("删除Pvclaimname失败" + err.Error())
	}
	return nil
}
