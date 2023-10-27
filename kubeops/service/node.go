package service

//列表
// 获取node 详情
import (
	"context"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

type node struct{}

var Node node

type Nodelist struct {
	Total    int           `json:"total"`
	NodeName []corev1.Node `json:"node_name"`
}

type NodeDetail struct {
	Detail          *corev1.Node `json:"detail"`
	Pods            []poddetail  `json:"pods"`
	MemoryAllocator string       `json:"memory_allocator"`
	AGE             string       `json:"age"`
	Total           int          `json:"total"`
}

type poddetail struct {
	Name    string
	Image   []string
	Labels  map[string]string
	Status  corev1.PodPhase
	Restart int32
	PodAge  string
}

// GetNodelist 列表
func (n *node) GetNodelist() (nodelist *Nodelist, err error) {
	nodelists, err := K8s.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取Node 失败" + err.Error())
		return nil, errors.New("获取Node 失败" + err.Error())
	}
	nodes := make([]corev1.Node, 0, len(nodelists.Items))
	for _, v := range nodelists.Items {
		nodes = append(nodes, v)
	}
	return &Nodelist{
		Total:    len(nodelists.Items),
		NodeName: nodes,
	}, err
}

// GetNodeDetail  获取namespace 详情
func (n *node) GetNodeDetail(NodeName string) (details *NodeDetail, err error) {
	//获取deploy
	detail, err := K8s.Clientset.CoreV1().Nodes().Get(context.TODO(), NodeName, metav1.GetOptions{})
	if err != nil {
		logger.Info("获取Node 详情失败" + err.Error())
		return nil, errors.New("获取Node 详情失败" + err.Error())
	}
	pods := n.GetNodePods(NodeName)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(detail.Status.Allocatable[corev1.ResourceRequestsCPU])
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	memory, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(detail.Status.Allocatable.Name(corev1.ResourceMemory, "").Value())/float64(1024*1024*1024)), 64)
	return &NodeDetail{
		Detail:          detail,
		Pods:            *pods,
		MemoryAllocator: strconv.FormatFloat(memory, 'f', -1, 64) + "Gi",
		AGE:             getAge(detail.CreationTimestamp.Unix()),
		Total:           len(*pods),
	}, nil
}

func (n *node) GetNodePods(NodeName string) (detail *[]poddetail) {
	//获取pod list
	podlist, err := K8s.Clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取Node 详情失败" + err.Error())
	}
	nodepods := make([]poddetail, 0)
	for _, pods := range podlist.Items {
		if pods.Spec.NodeName == NodeName {
			containers := make([]string, 0)
			for _, container := range pods.Spec.Containers {
				containers = append(containers, container.Image)
			}
			nodepods = append(nodepods, poddetail{
				Name:    pods.Name,
				Image:   containers,
				Labels:  pods.Labels,
				Status:  pods.Status.Phase,
				Restart: pods.Status.ContainerStatuses[0].RestartCount,
				PodAge:  getAge(pods.CreationTimestamp.Unix()),
			})
		}
	}
	return &nodepods
}

// getAge 获取node 时间戳
func getAge(CTT int64) string {
	age := time.Now().Unix() - CTT
	switch {
	case age < 60:
		//秒
		return strconv.FormatInt(age, 10) + "s"
	case age < 3600:
		//分
		return strconv.FormatInt(age/60, 10) + "min"
	case age < 86400:
		//时
		return strconv.FormatInt(age/3600, 10) + "h"
	case age < 31536000:
		//天
		return strconv.FormatInt(age/86400, 10) + "d"
	default:
		//年
		return strconv.FormatInt(age/31536000, 10) + "y"
	}

}
