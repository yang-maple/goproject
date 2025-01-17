package service

import (
	"context"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeops/model"
)

type pvClaim struct{}
type PvClaimResp struct {
	Total int       `json:"total"`
	Item  []pvcInfo `json:"item"`
}
type pvcInfo struct {
	Name         string                            `json:"name"`
	Namespace    string                            `json:"namespace"`
	Labels       map[string]string                 `json:"labels"`
	Status       corev1.PersistentVolumeClaimPhase `json:"status"`
	Volume       string                            `json:"volume"`
	Claim        string                            `json:"claim"`
	AccessModes  []string                          `json:"access_modes"`
	StorageClass string                            `json:"storage_class"`
	Age          string                            `json:"age"`
}

type CreateClaim struct {
	Name             string                              `json:"name"`
	Namespace        string                              `json:"namespace"`
	Labels           map[string]string                   `json:"labels"`
	AccessModes      []corev1.PersistentVolumeAccessMode `json:"access_modes"`
	Storage          string                              `json:"storage"`
	StorageClassName string                              `json:"storage_classname" default:"nfs-client"`
}

var Claim pvClaim

func (p *pvClaim) toCells(pvClaims []corev1.PersistentVolumeClaim) []DataCell {
	cells := make([]DataCell, len(pvClaims))
	for i := range pvClaims {
		cells[i] = pvcCell(pvClaims[i])
	}
	return cells
}

func (p *pvClaim) fromCells(cells []DataCell) []corev1.PersistentVolumeClaim {
	pvClaims := make([]corev1.PersistentVolumeClaim, len(cells))
	for i := range cells {
		pvClaims[i] = corev1.PersistentVolumeClaim(cells[i].(pvcCell))
	}
	return pvClaims
}

// GetPVClaimList 列表
func (p *pvClaim) GetPVClaimList(claimName, Namespace string, Limit, Page int) (DP *PvClaimResp, err error) {
	//获取deployment 的所有清单列表
	claimList, err := K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 PersistentVolumeClaims 失败" + err.Error())
		return nil, err
	}
	//组装数据
	selectData := &dataselector{
		GenericDataList: p.toCells(claimList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{claimName},
			Paginate: &PaginateQuery{
				limit: Limit,
				page:  Page,
			},
		},
	}
	//先过滤 后排序
	filtered := selectData.Filter()
	total := len(filtered.GenericDataList)
	//分页
	dataPage := filtered.Sort().Pagination()
	pvClaims := p.fromCells(dataPage.GenericDataList)
	item := make([]pvcInfo, 0, total)
	for _, v := range pvClaims {
		aModes := make([]string, 0)
		for _, value := range v.Spec.AccessModes {
			aModes = append(aModes, string(value))
		}
		item = append(item, pvcInfo{
			Name:         v.Name,
			Namespace:    v.Namespace,
			Labels:       v.Labels,
			Status:       v.Status.Phase,
			Volume:       v.Spec.VolumeName,
			Claim:        v.Spec.Resources.Requests.Storage().String(),
			AccessModes:  aModes,
			StorageClass: *v.Spec.StorageClassName,
			Age:          model.GetAge(v.CreationTimestamp.Unix()),
		})
	}
	return &PvClaimResp{
		Total: total,
		Item:  item,
	}, nil
}

// GetPVClaimDetail 详情
func (p *pvClaim) GetPVClaimDetail(Namespace, claimName string) (detail *corev1.PersistentVolumeClaim, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).Get(context.TODO(), claimName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取 PersistentVolumeClaims 详情失败" + err.Error())
		return nil, err
	}
	detail.Kind = "PersistentVolumeClaim"
	detail.APIVersion = "v1"
	return detail, nil
}

// CreatePVClaim 创建
func (p *pvClaim) CreatePVClaim(data *CreateClaim) (err error) {

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: data.AccessModes,
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
		logger.Info("创建 PersistentVolumeClaims 失败" + err.Error())
		return err
	}
	return nil

}

// DelPVClaim 删除
func (p *pvClaim) DelPVClaim(Namespace, claimName string) (err error) {
	err = K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).Delete(context.TODO(), claimName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除PersistentVolumeClaim失败" + err.Error())
		return err
	}
	return nil
}

// UpdatePVClaim 更新
func (p *pvClaim) UpdatePVClaim(Namespace string, claimName *corev1.PersistentVolumeClaim) (err error) {
	_, err = K8s.Clientset.CoreV1().PersistentVolumeClaims(Namespace).Update(context.TODO(), claimName, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("更新 PersistentVolumeClaim 失败" + err.Error())
		return err
	}
	return nil
}
