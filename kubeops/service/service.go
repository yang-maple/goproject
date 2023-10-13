package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type service struct{}

var Services service

type Svcresp struct {
	Total int              `json:"total"`
	Item  []corev1.Service `json:"item"`
}

type CreateService struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Labels       map[string]string `json:"labels"`
	Type         string            `json:"type"`
	Serviceports []Serviceport     `json:"service_port"`
}
type Serviceport struct {
	Portname   string          `json:"port_name"`
	Port       int32           `json:"port"`
	Protocol   corev1.Protocol `json:"protocol"`
	TargetPort int32           `json:"target_port"`
	NodePort   int32           `json:"node_port"`
}

// 数据类型转换 tocells
func (s *service) tocells(svc []corev1.Service) []DataCell {
	cells := make([]DataCell, len(svc))
	for i := range svc {
		cells[i] = serviceCell(svc[i])
	}
	return cells
}

func (s *service) fromcells(cells []DataCell) []corev1.Service {
	svc := make([]corev1.Service, len(cells))
	for i := range cells {
		svc[i] = corev1.Service(cells[i].(serviceCell))
	}
	return svc
}

// GetSvcList 获取services列表
func (s *service) GetSvcList(Svcname, Namespace string, Limit, Page int) (DP *Svcresp, err error) {
	//获取deployment 的所有清单列表
	svclist, err := K8s.Clientset.CoreV1().Services(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取 svclist 失败" + err.Error())
	}
	//组装数据
	selectdata := &dataselector{
		GenericDataList: s.tocells(svclist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Svcname},
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
	svcs := s.fromcells(datapage.GenericDataList)
	for _, v := range svcs {
		fmt.Println(v.Name, v.CreationTimestamp.Time)
	}
	return &Svcresp{
		Total: total,
		Item:  svcs,
	}, err
}

// GetSvcDetail 获取 services 详情
func (s *service) GetSvcDetail(Namespace, Svcname string) (detail *corev1.Service, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().Services(Namespace).Get(context.TODO(), Svcname, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取services 详情失败" + err.Error())
		return nil, errors.New("获取services 详情失败" + err.Error())
	}
	return detail, nil
}

// CreateSvc 创建 services
func (s *service) CreateSvc(data *CreateService) (err error) {
	fmt.Println(data)
	createsvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   data.Name,
			Labels: data.Labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: data.Labels,
			Type:     "",
		},
		Status: corev1.ServiceStatus{},
	}
	for i := range data.Serviceports {
		specport := corev1.ServicePort{
			Name:     data.Serviceports[i].Portname,
			Protocol: data.Serviceports[i].Protocol,
			Port:     data.Serviceports[i].Port,
			TargetPort: intstr.IntOrString{
				Type:   0,
				IntVal: data.Serviceports[i].TargetPort,
			},
			NodePort: 0,
		}
		if data.Type == "NodePort" && data.Serviceports[i].NodePort != 0 {
			createsvc.Spec.Type = "NodePort"
			specport.NodePort = data.Serviceports[i].NodePort
			fmt.Println(i, data.Serviceports[i].NodePort)
		}
		createsvc.Spec.Ports = append(createsvc.Spec.Ports, specport)
	}
	fmt.Println(createsvc)
	_, err = K8s.Clientset.CoreV1().Services(data.Namespace).Create(context.TODO(), createsvc, metav1.CreateOptions{})
	if err != nil {
		logger.Info("创建 service 失败" + err.Error())
		return errors.New("创建 service 失败" + err.Error())
	}
	return nil

}

// DelSvc 删除 services
func (s *service) DelSvc(Namespace, Svcname string) (err error) {
	err = K8s.Clientset.CoreV1().Services(Namespace).Delete(context.TODO(), Svcname, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除service失败" + err.Error())
		return errors.New("删除service失败" + err.Error())
	}
	return nil
}

// UpdateSvc 更新 services
func (s *service) UpdateSvc(Namespace string, svc *corev1.Service) (err error) {
	_, err = K8s.Clientset.CoreV1().Services(Namespace).Update(context.TODO(), svc, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("service 更新失败" + err.Error())
		return errors.New("service 更新失败" + err.Error())
	}
	return nil
}
