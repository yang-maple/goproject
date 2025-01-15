package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
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
	Createdeployment(clientset)
}

func CreateNamespace(clientset *kubernetes.Clientset, newnamespace string) {
	var name = &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: newnamespace,
		},
	}
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), name, metav1.CreateOptions{})
	if err == nil {
		fmt.Println("创建成功")
	} else {
		fmt.Println("创建失败", err)
	}
}

func Createdeployment(clientset *kubernetes.Clientset) {
	var deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := clientset.AppsV1().Deployments("jenkins").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err == nil {
		fmt.Println("创建成功")
	} else {
		fmt.Println(err)
	}
}

func Int32Ptr(i int32) *int32 { return &i }
