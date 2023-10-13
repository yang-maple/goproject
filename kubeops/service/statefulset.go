package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 定义空结构体
type statefulset struct{}

// 全局变量供外部调用
var Statefulset statefulset

// 定义返回的结构体
type StsResp struct {
	Total int                  `json:"total"`
	Sts   []appsv1.StatefulSet `json:"sts"`
}

// 将statefulsetcell 转换为 datacell
func (s statefulset) tocells(statecell []appsv1.StatefulSet) []DataCell {
	cells := make([]DataCell, len(statecell))
	for i := range cells {
		cells[i] = statefulsetCell(statecell[i])
	}
	return cells
}

// 将datacell 转换为 statefulsetcell
func (s statefulset) fromcells(cells []DataCell) []appsv1.StatefulSet {
	statefulcell := make([]appsv1.StatefulSet, len(cells))
	for i := range statefulcell {
		statefulcell[i] = appsv1.StatefulSet(cells[i].(statefulsetCell))
	}
	return statefulcell
}

// 列表
func (s *statefulset) GetStsList(StsName, Namespace string, Limit, Page int) (stsresp *StsResp, err error) {
	stslist, err := K8s.Clientset.AppsV1().StatefulSets(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取stslist 失败" + err.Error())
		return nil, errors.New("获取stslist 失败" + err.Error())
	}
	selectdata := &dataselector{
		GenericDataList: s.tocells(stslist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{StsName},
			Paginate: &PaginateQuery{
				limit: Limit,
				page:  Page,
			},
		},
	}

	selectdata.Filter()
	//先过滤 后排序
	filtered := selectdata.Filter()
	total := len(filtered.GenericDataList)
	//分页
	datapage := filtered.Sort().Pagination()
	states := s.fromcells(datapage.GenericDataList)
	for _, v := range states {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &StsResp{
		Total: total,
		Sts:   states,
	}, err

}

// 详情
func (s *statefulset) GetStsDetal(Namespace, StsName string) (detail *appsv1.StatefulSet, err error) {
	//获取deploy
	detail, err = K8s.Clientset.AppsV1().StatefulSets(Namespace).Get(context.TODO(), StsName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取statefulset 详情失败" + err.Error())
		return nil, errors.New("获取statefulset 详情失败" + err.Error())
	}
	return detail, nil
}

// 删除
func (s *statefulset) DelSts(Namespace, StsName string) (err error) {
	err = K8s.Clientset.AppsV1().StatefulSets(Namespace).Delete(context.TODO(), StsName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除实例失败" + err.Error())
		return errors.New("删除实例失败" + err.Error())
	}
	return nil
}

// 更新
func (s *statefulset) UpdataSts(Namespace string, Sts *appsv1.StatefulSet) (err error) {
	_, err = K8s.Clientset.AppsV1().StatefulSets(Namespace).Update(context.TODO(), Sts, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("StatefulSets 更新失败" + err.Error())
		return errors.New("StatefulSets 更新失败" + err.Error())
	}

	return nil
}
