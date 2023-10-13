package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type configmap struct{}
type Cmsresp struct {
	Total int                `json:"total"`
	Item  []corev1.ConfigMap `json:"item"`
}

type Createconfigmap struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
	Data      map[string]string `json:"data"`
	//是否能进行修改的标识符
	Ismodity bool `json:"ismodity"`
}

var Configmaps configmap

func (c configmap) tocells(cms []corev1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(cms))
	for i := range cms {
		cells[i] = configmapCell(cms[i])
	}
	return cells
}

func (c configmap) fromcells(cells []DataCell) []corev1.ConfigMap {
	cms := make([]corev1.ConfigMap, len(cells))
	for i := range cells {
		cms[i] = corev1.ConfigMap(cells[i].(configmapCell))
	}
	return cms
}

// 列表
func (c *configmap) GetCmList(Configname, Namespace string, Limit, Page int) (DP *Cmsresp, err error) {
	//获取deployment 的所有清单列表
	cmlist, err := K8s.Clientset.CoreV1().ConfigMaps(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 configmap 失败" + err.Error())
	}

	//组装数据
	selectdata := &dataselector{
		GenericDataList: c.tocells(cmlist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Configname},
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
	cms := c.fromcells(datapage.GenericDataList)
	for _, v := range cms {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &Cmsresp{
		Total: total,
		Item:  cms,
	}, err
}

// 详情
func (c *configmap) GetCmDetail(Namespace, Configname string) (detail *corev1.ConfigMap, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().ConfigMaps(Namespace).Get(context.TODO(), Configname, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取configmap 详情失败" + err.Error())
		return nil, errors.New("获取configmap 详情失败" + err.Error())
	}
	return detail, nil
}

// 创建
func (c *configmap) CreateCm(data *Createconfigmap) (err error) {
	config := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Immutable: &data.Ismodity,
		Data:      data.Data,
	}
	_, err = K8s.Clientset.CoreV1().ConfigMaps(data.Namespace).Create(context.TODO(), config, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 configmap 失败" + err.Error())
		return errors.New("创建 configmap 失败" + err.Error())
	}
	return nil

}

// 删除
func (c *configmap) DelCm(Namespace, Configname string) (err error) {
	err = K8s.Clientset.CoreV1().ConfigMaps(Namespace).Delete(context.TODO(), Configname, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除configmap失败" + err.Error())
		return errors.New("删除configmap失败" + err.Error())
	}
	return nil
}

// 更新
func (c *configmap) UpdataCm(Namespace string, Config *corev1.ConfigMap) (err error) {
	_, err = K8s.Clientset.CoreV1().ConfigMaps(Namespace).Update(context.TODO(), Config, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("configmap更新失败" + err.Error())
		return errors.New("configmap 更新失败" + err.Error())
	}
	return nil
}
