package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
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
}
