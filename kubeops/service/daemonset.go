package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeops/model"
)

var DaemonSet daemonSet

type daemonSet struct{}

type DaemonDp struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type DsResp struct {
	Total int          `json:"total"`
	Item  []daemonInfo `json:"item"`
}

type daemonInfo struct {
	Name       string            `json:"name"`
	Namespaces string            `json:"namespaces"`
	Image      []string          `json:"image"`
	Labels     map[string]string `json:"labels"`
	Pods       string            `json:"pods"`
	Age        string            `json:"age"`
	Status     string            `json:"status"`
}

// toCell 数据类型转换 from daemonCell to dataCell
func (d *daemonSet) toCell(daemonSets []appsv1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(daemonSets))
	for i := range daemonSets {
		cells[i] = daemonsetCell(daemonSets[i])
	}
	return cells
}

// fromCell 数据类型转换 from dataCell to daemonCell
func (d *daemonSet) fromCell(cells []DataCell) []appsv1.DaemonSet {
	daemonSets := make([]appsv1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsv1.DaemonSet(cells[i].(daemonsetCell))
	}
	return daemonSets
}

// GetDsList 列表
func (d *daemonSet) GetDsList(DsName, Namespace string, Limit, Page int) (DP *DsResp, err error) {
	//获取deployment 的所有清单列表
	daemonList, err := K8s.Clientset.AppsV1().DaemonSets(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 daemonSetList 失败" + err.Error())
	}
	//组装数据
	selectData := &dataselector{
		GenericDataList: d.toCell(daemonList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{DsName},
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
	dss := d.fromCell(dataPage.GenericDataList)
	item := make([]daemonInfo, 0, total)
	for _, v := range dss {
		images := make([]string, 0, len(v.Spec.Template.Spec.Containers))
		for _, im := range v.Spec.Template.Spec.Containers {
			images = append(images, im.Image)
		}
		pods, status := model.GetStatus(v.Status.DesiredNumberScheduled, v.Status.NumberReady)
		item = append(item, daemonInfo{
			Name:       v.Name,
			Namespaces: v.Namespace,
			Image:      images,
			Labels:     v.Labels,
			Pods:       pods,
			Age:        model.GetAge(v.CreationTimestamp.Unix()),
			Status:     status,
		})
	}
	return &DsResp{
		Total: total,
		Item:  item,
	}, err
}

// GetDsDetail  获取详情
func (d *daemonSet) GetDsDetail(Namespace, DsName string) (detail *appsv1.DaemonSet, err error) {
	//获取deploy
	detail, err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Get(context.TODO(), DsName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取deployment 详情失败" + err.Error())
		return nil, errors.New("获取deployment 详情失败" + err.Error())
	}
	detail.Kind = "DaemonSet"
	detail.APIVersion = "apps/v1"
	return detail, nil
}

// DelDs 删除
func (d *daemonSet) DelDs(Namespace, DsName string) (err error) {
	err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Delete(context.TODO(), DsName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除实例失败" + err.Error())
		return errors.New("删除实例失败" + err.Error())
	}
	return nil
}

// UpdateDs 更新
func (d *daemonSet) UpdateDs(Namespace string, ds *appsv1.DaemonSet) (err error) {
	_, err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Update(context.TODO(), ds, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("daemonSet 更新失败" + err.Error())
		return errors.New("daemonSet 更新失败" + err.Error())
	}

	return nil
}
