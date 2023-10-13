package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Daemonset daemonset

type daemonset struct{}

type Dsdp struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type DsResp struct {
	Total int                `json:"total"`
	Item  []appsv1.DaemonSet `json:"item"`
}

// 数据类型转换
// daemonsetcell to datacell

func (d *daemonset) todatacell(daemonsets []appsv1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(daemonsets))
	for i := range daemonsets {
		cells[i] = daemonsetCell(daemonsets[i])
	}
	return cells
}

// datacell to daemonsetcell
func (d *daemonset) fromdatacell(cells []DataCell) []appsv1.DaemonSet {
	daemonsets := make([]appsv1.DaemonSet, len(cells))
	for i := range cells {
		daemonsets[i] = appsv1.DaemonSet(cells[i].(daemonsetCell))
	}
	return daemonsets
}

// 列表
func (d *daemonset) GetDsList(DsName, Namespace string, Limit, Page int) (DP *DsResp, err error) {
	//获取deployment 的所有清单列表
	dslist, err := K8s.Clientset.AppsV1().DaemonSets(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取daemonsetlist 失败" + err.Error())
	}
	//组装数据
	selectdata := &dataselector{
		GenericDataList: d.todatacell(dslist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{DsName},
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
	dss := d.fromdatacell(datapage.GenericDataList)
	for _, v := range dss {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &DsResp{
		Total: total,
		Item:  dss,
	}, err
}

// 获取详情
func (d *daemonset) GetDsDetal(Namespace, DsName string) (detail *appsv1.DaemonSet, err error) {
	//获取deploy
	detail, err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Get(context.TODO(), DsName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取deployment 详情失败" + err.Error())
		return nil, errors.New("获取deployment 详情失败" + err.Error())
	}
	return detail, nil
}

// 删除
func (d *daemonset) DelDs(Namespace, DsName string) (err error) {
	err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Delete(context.TODO(), DsName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除实例失败" + err.Error())
		return errors.New("删除实例失败" + err.Error())
	}
	return nil
}

// 更新
func (d *daemonset) UpdataDs(Namespace string, ds *appsv1.DaemonSet) (err error) {
	_, err = K8s.Clientset.AppsV1().DaemonSets(Namespace).Update(context.TODO(), ds, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("daemonset 更新失败" + err.Error())
		return errors.New("daemonset 更新失败" + err.Error())
	}

	return nil
}
