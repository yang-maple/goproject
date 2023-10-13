package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type namespace struct{}

var Namespace namespace

type Nslist struct {
	Total int      `json:"total"`
	Item  []string `json:"item"`
}

// 列表
func (n *namespace) GetNslist() (Namespaceslist *Nslist, err error) {
	nslist, err := K8s.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取namespaceslist 失败" + err.Error())
		return nil, errors.New("获取namespaceslist 失败" + err.Error())
	}
	list := make([]string, 0, len(nslist.Items))
	for _, v := range nslist.Items {
		list = append(list, v.Name)
	}
	return &Nslist{
		Total: len(nslist.Items),
		Item:  list,
	}, err
}

// 获取namespace 详情
func (n *namespace) GetNsDetal(NsName string) (detail *corev1.Namespace, err error) {
	//获取deploy
	detail, err = K8s.Clientset.CoreV1().Namespaces().Get(context.TODO(), NsName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取Namespace 详情失败" + err.Error())
		return nil, errors.New("获取Namespace 详情失败" + err.Error())
	}
	return detail, nil
}

// 删除
func (n *namespace) DelNs(NsName string) (err error) {
	err = K8s.Clientset.CoreV1().Namespaces().Delete(context.TODO(), NsName, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除实例失败" + err.Error())
		return errors.New("删除实例失败" + err.Error())
	}
	return nil
}

// 创建
func (n *namespace) CreateNs(NsName string) (err error) {
	namespaceconfig := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NsName,
		},
	}
	_, err = K8s.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespaceconfig, metav1.CreateOptions{})
	if err != nil {
		logger.Info("namespace 创建失败" + err.Error())
		return errors.New("namespace 创建失败" + err.Error())
	}

	return nil
}
