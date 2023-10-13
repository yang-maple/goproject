package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ingress struct{}

var Ingress ingress

type Ingresp struct {
	Total int                 `json:"total"`
	Item  []networkv1.Ingress `json:"item"`
}

type Createingress struct {
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	Labels           map[string]string `json:"labels"`
	IngressClassName string            `json:"ingress_class_name"`
	Rules            []Rule            `json:"rules"`
}
type Rule struct {
	Host                  string     `json:"host"`
	HTTPIngressRuleValues []HTTPRule `json:"http_ingress_rule_values"`
}
type HTTPRule struct {
	Path        string `json:"path"`
	ServiceName string `json:"service_name"`
	ServicePort int32  `json:"service_port"`
}

func (i *ingress) tocells(ing []networkv1.Ingress) []DataCell {
	cells := make([]DataCell, len(ing))
	for i := range ing {
		cells[i] = ingressCell(ing[i])
	}
	return cells
}

func (i *ingress) fromcells(cells []DataCell) []networkv1.Ingress {
	svc := make([]networkv1.Ingress, len(cells))
	for i := range cells {
		svc[i] = networkv1.Ingress(cells[i].(ingressCell))
	}
	return svc
}

// GetIngList 获取 ingress 资源列表
func (i *ingress) GetIngList(Ingname, Namespace string, Limit, Page int) (DP *Ingresp, err error) {
	//获取deployment 的所有清单列表
	inglist, err := K8s.Clientset.NetworkingV1().Ingresses(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 ingress 失败" + err.Error())
	}

	//组装数据
	selectdata := &dataselector{
		GenericDataList: i.tocells(inglist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Ingname},
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
	ings := i.fromcells(datapage.GenericDataList)
	for _, v := range ings {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &Ingresp{
		Total: total,
		Item:  ings,
	}, err
}

// GetIngDetail 获取 ingress 资源详情
func (i *ingress) GetIngDetail(Namespace, Ingname string) (detail *networkv1.Ingress, err error) {
	//获取deploy
	detail, err = K8s.Clientset.NetworkingV1().Ingresses(Namespace).Get(context.TODO(), Ingname, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取ingress 详情失败" + err.Error())
		return nil, errors.New("获取ingress 详情失败" + err.Error())
	}
	return detail, nil
}

// CreateIng 创建 ingress 资源
func (i *ingress) CreateIng(data *Createingress) (err error) {
	var pathtype networkv1.PathType = networkv1.PathTypeImplementationSpecific
	createing := &networkv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Spec: networkv1.IngressSpec{
			IngressClassName: &data.IngressClassName,
			TLS:              nil,
		},
		Status: networkv1.IngressStatus{},
	}
	for i := range data.Rules {
		specrole := networkv1.IngressRule{
			Host: data.Rules[i].Host,
		}
		for j := range data.Rules[i].HTTPIngressRuleValues {
			specroleshttppath := &networkv1.HTTPIngressRuleValue{[]networkv1.HTTPIngressPath{
				{
					Path:     data.Rules[i].HTTPIngressRuleValues[j].Path,
					PathType: &pathtype,
					Backend: networkv1.IngressBackend{
						Service: &networkv1.IngressServiceBackend{
							Name: data.Rules[i].HTTPIngressRuleValues[j].ServiceName,
							Port: networkv1.ServiceBackendPort{
								Name:   "",
								Number: data.Rules[i].HTTPIngressRuleValues[j].ServicePort,
							},
						},
					},
				},
			}}
			specrole.IngressRuleValue.HTTP = specroleshttppath
		}
		createing.Spec.Rules = append(createing.Spec.Rules, specrole)
	}
	_, err = K8s.Clientset.NetworkingV1().Ingresses(data.Namespace).Create(context.TODO(), createing, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 service 失败" + err.Error())
		return errors.New("创建 service 失败" + err.Error())
	}
	return nil

}

// DelIng 删除ingress 资源
func (i *ingress) DelIng(Namespace, Ingname string) (err error) {
	err = K8s.Clientset.NetworkingV1().Ingresses(Namespace).Delete(context.TODO(), Ingname, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除ingress失败" + err.Error())
		return errors.New("删除ingress失败" + err.Error())
	}
	return nil
}

// UpdateIng 更新ingress 资源详情
func (i *ingress) UpdateIng(Namespace string, ing *networkv1.Ingress) (err error) {
	_, err = K8s.Clientset.NetworkingV1().Ingresses(Namespace).Update(context.TODO(), ing, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("ingress 更新失败" + err.Error())
		return errors.New("ingress 更新失败" + err.Error())
	}
	return nil
}
