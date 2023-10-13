package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"kubeops/config"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Pod pod

type pod struct {
}
type Podnp struct {
	Name   string
	Number int
}

//类型转换的两个方法
// corev1.pod --> Datacell

func (p *pod) toCells(pods []corev1.Pod) []DataCell {
	// 定义并初始化 []DataCell 数组
	cells := make([]DataCell, len(pods))
	//数据转换，将pods类型转化为datacells 类型
	for i := range pods {
		cells[i] = podCell(pods[i])
	}
	return cells
}

// Datacell --> corev1.pod
func (p *pod) fromCells(cells []DataCell) []corev1.Pod {

	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		//先断言判断是否为 podcell 然后进行转换
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}

type PodResp struct {
	Total int          `json:"total"`
	Item  []corev1.Pod `json:"item"`
}

// 定义 getpod 方法获取pod 列表 支持过滤排序和分页
func (p *pod) GetPods(FilterName, NameSpaces string, Limit, Page int) (PR *PodResp, err error) {
	//获取postlist
	Postlist, err := K8s.Clientset.CoreV1().Pods(NameSpaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取postlist 失败" + err.Error())
	}
	//实例化dataselect对象
	selectdata := &dataselector{
		GenericDataList: p.toCells(Postlist.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{FilterName},
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
	pods := p.fromCells(datapage.GenericDataList)
	for _, v := range pods {
		fmt.Println(v.Name, v.Namespace, v.CreationTimestamp.Time)
	}
	return &PodResp{
		Total: total,
		Item:  pods,
	}, err
}

// 获取pod 详情
func (p *pod) GetPodDetail(Name, Namespace string) (pod *corev1.Pod, err error) {
	poddetail, err := K8s.Clientset.CoreV1().Pods(Namespace).Get(context.TODO(), Name, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取pod 详情失败" + err.Error())
		return nil, errors.New("获取pod 详情失败")
	}
	return poddetail, nil
}

// 删除 pod
func (p *pod) DelPod(Name, Namespace string) error {
	err := K8s.Clientset.CoreV1().Pods(Namespace).Delete(context.TODO(), Name, metav1.DeleteOptions{})
	if err != nil {
		logger.Info("删除失败" + err.Error())
		return errors.New("删除失败" + err.Error())
	}
	return nil
}

// 更新 pod
func (p *pod) UpdataPod(npod *corev1.Pod, Namespaces string) error {
	_, err := K8s.Clientset.CoreV1().Pods(Namespaces).Update(context.TODO(), npod, metav1.UpdateOptions{})
	if err != nil {
		logger.Info("更新失败" + err.Error())
		return errors.New("更新失败" + err.Error())
	}
	return nil
}

// 获取pod 的容器列表
func (p *pod) GetContainer(podName, Namespaces string) (containers []string, err error) {
	pod, err := p.GetPodDetail(podName, Namespaces)
	if err != nil {
		logger.Info("获取pod 失败" + err.Error())
		return nil, errors.New("获取pod 失败")
	}
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

// 获取pod 内容器日志
func (p *pod) GetContainerLog(podname, containername, namespaces string) (log string, err error) {
	taillines := int64(config.Logtaillimit)
	//设置日志的配置，容器名，获取的内容的配置
	logopt := &corev1.PodLogOptions{
		Container: containername,
		Follow:    false,
		TailLines: &taillines,
	}

	//获取一个request 实例
	request := K8s.Clientset.CoreV1().Pods(namespaces).GetLogs(podname, logopt)

	//发起stream连接，到底response.body
	podlog, err := request.Stream(context.TODO())
	if err != nil {
		logger.Info("获取podlog 失败" + err.Error())
		return "", errors.New("获取podlog 失败")
	}
	//延时关闭
	defer podlog.Close()
	//将response.body 写入缓冲区，用于转化成string 类型
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podlog)
	if err != nil {
		logger.Info("写入缓冲失败 失败" + err.Error())
		return "", errors.New("写入缓冲失败 失败")
	}
	return buf.String(), nil
}

// 获取每个namespace 下pod的数量
func (p *pod) Countpod() (total []Podnp, err error) {
	//获取所有的namespaces
	namespacelist, err := K8s.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取namespace 列表失败" + err.Error())
		return nil, errors.New("获取namespace 列表失败")
	}

	////获取每个ns下的pod数量
	for _, namespace := range namespacelist.Items {
		podlist, err := K8s.Clientset.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("获取namespace下podlist 列表失败" + err.Error())
			return nil, errors.New("获取namespace下podlist 列表失败")
		}
		//组装数据
		npm := &Podnp{
			Name:   namespace.Name,
			Number: len(podlist.Items),
		}
		fmt.Println(npm)
		total = append(total, *npm)
	}
	fmt.Println(total)
	return total, nil
}
