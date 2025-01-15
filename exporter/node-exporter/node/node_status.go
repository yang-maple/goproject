package node

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

var publicdir = "../config"

type nodeStatus struct {
	desc *prometheus.Desc
}

func NewNodeStatusController() *nodeStatus {
	desc := prometheus.NewDesc(prometheus.BuildFQName("K8S", "NODE", "STATUS"), "capability", []string{"Config_Name", "Node"}, nil)
	return &nodeStatus{desc}
}

func (c *nodeStatus) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *nodeStatus) Collect(metrics chan<- prometheus.Metric) {
	//cmds := []string{"cores", "used", "load1", "load5", "load15"}
	files, err := os.ReadDir(publicdir)
	if err != nil {
		fmt.Println("Failed For read file", err)
	}
	for _, file := range files {
		info := get_node(file.Name())
		for node, status := range info {
			metrics <- prometheus.MustNewConstMetric(
				c.desc,
				prometheus.GaugeValue,
				status,
				file.Name(),
				node,
			)
		}
	}
}

func get_node(name string) map[string]float64 {
	config, err := clientcmd.BuildConfigFromFlags("", publicdir+"/"+name)
	if err != nil {
		log.Fatal("config err :", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("clientset err :", err)
	}
	nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("nodelist err :", err)
	}
	list := make(map[string]float64)
	for _, node := range nodelist.Items {
		//获取每个节点的状态，根据状态返回值
		for _, status := range node.Status.Conditions {
			if status.Type == "Ready" && status.Status == "True" {
				list[node.Name] = 1
				//结束Conditions的循环，继续节点的循环
				break
			}
			list[node.Name] = 0
		}
	}
	return list
}
