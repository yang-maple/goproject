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
	nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("nodelist err :", err)
	}
	for _, node := range nodelist.Items {
		fmt.Println(node.Name, node.Status.Addresses[0].Address)
	}

	namespacelist, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("namespacelist err :", err)
	}
	for _, namespaces := range namespacelist.Items {
		fmt.Println(namespaces.Name)
	}

}
