package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
<<<<<<< HEAD
=======
	"strconv"
>>>>>>> 4622a0a (2025-1-15)
)

func main() {
	kubeconfig := "config/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("config err :", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("clientset err :", err)
	}
<<<<<<< HEAD

	//创建 namespace
	//var newnamespace = "jenkins"
	//CreateNamespace(clientset, newnamespace)

	//创建 deployment
	UpdateDeployment(clientset)
}

func UpdateDeployment(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments("jenkins")
	result, getErr := deploymentsClient.Get(context.TODO(), "nginx", metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}
	result.Spec.Replicas = Int32Ptr(4)                           // reduce replica count
	result.Spec.Template.Spec.Containers[0].Image = "nginx:1.16" // change nginx version
	_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	if updateErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", updateErr))
	}
=======
	pods, _ := clientset.CoreV1().Pods(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	cpulimit := make(map[string]float64)
	cpurequest := make(map[string]float64)
	memlimit := make(map[string]int64)
	memrequest := make(map[string]int64)
	for _, pod := range pods.Items {
		if pod.Spec.NodeName != "" {
			for _, container := range pod.Spec.Containers {
				if _, ok := cpulimit[pod.Spec.NodeName]; ok {
					cpulimit[pod.Spec.NodeName] += float62(container.Resources.Limits.Cpu().AsApproximateFloat64())
					cpurequest[pod.Spec.NodeName] += float62(container.Resources.Requests.Cpu().AsApproximateFloat64())
					memlimit[pod.Spec.NodeName] += container.Resources.Limits.Memory().Value()
					memrequest[pod.Spec.NodeName] += container.Resources.Requests.Memory().Value()
					continue
				} else {
					cpulimit[pod.Spec.NodeName] += float62(container.Resources.Limits.Cpu().AsApproximateFloat64())
					cpurequest[pod.Spec.NodeName] += float62(container.Resources.Requests.Cpu().AsApproximateFloat64())
					memlimit[pod.Spec.NodeName] += container.Resources.Limits.Memory().Value()
					memrequest[pod.Spec.NodeName] += container.Resources.Requests.Memory().Value()
				}
			}
		}
	}
	fmt.Println(cpulimit)
	fmt.Println(cpurequest)
	fmt.Println(memlimit)
	fmt.Println(memrequest)
}

func float62(v float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v*1000), 64)
	return value
>>>>>>> 4622a0a (2025-1-15)
}
