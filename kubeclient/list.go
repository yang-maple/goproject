package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
<<<<<<< HEAD
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
=======
	"k8s.io/client-go/kubernetes"
>>>>>>> 4622a0a (2025-1-15)
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	kubeconfig := "config/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("config err :", err)
	}
<<<<<<< HEAD
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	log.Fatal("clientset err :", err)
	//}

	//nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	log.Fatal("nodelist err :", err)
	//}
	//for _, node := range nodelist.Items {
	//	fmt.Println(node.Name, node.Status.Addresses[0].Address)
	//}
=======
	// 证书检查
	//fmt.Println(string(config.KeyData))
	//block, _ := pem.Decode(config.CertData)
	//if block == nil {
	//	log.Fatal("failed to parse certificate PEM")
	//}
	//cert, err := x509.ParseCertificate(block.Bytes)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Issuer: %s\n", cert.Issuer)
	//fmt.Printf("Subject: %s\n", cert.Subject)
	//fmt.Printf("Valid from: %s to %s\n", cert.NotBefore, cert.NotAfter)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("clientset err :", err)
	}
	nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("nodelist err :", err)
	}
	for _, node := range nodelist.Items {
		fmt.Println(node.Status)
	}
	fmt.Println(nodelist.Items[1].Name)
	//fmt.Println(nodelist.Items[1].Status.Conditions)
>>>>>>> 4622a0a (2025-1-15)

	//namespacelist, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	log.Fatal("namespacelist err :", err)
	//}
	//for _, namespaces := range namespacelist.Items {
	//	fmt.Println(namespaces.Name)
	//}
	//deploy, _ := clientset.AppsV1().Deployments("jenkins").Get(context.TODO(), "nginx", metav1.GetOptions{})
	//
	//println(deploy.Annotations["deployment.kubernetes.io/revision"])
<<<<<<< HEAD
	client, err := dynamic.NewForConfig(config)
	gvr := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}
	deployment, err := client.Resource(gvr).Namespace("jenkins").Get(context.TODO(), "nginx", metav1.GetOptions{})
	currentRevision := deployment.GetAnnotations()["deployment.kubernetes.io/revision"]
	fmt.Println("current revision:", currentRevision)
	historyRevision := "5"
	patchData := fmt.Sprintf(`{"metadata":{"annotations":{"deployment.kubernetes.io/revision":"%s"}}}`, historyRevision)
	if _, err = client.Resource(gvr).Namespace("jenkins").Patch(context.TODO(), "nginx", types.StrategicMergePatchType, []byte(patchData), metav1.PatchOptions{}); err != nil {
		log.Println("Patch", err)
		return
	}

	fmt.Println("rollback done")
=======
	//client, err := dynamic.NewForConfig(config)
	//gvr := schema.GroupVersionResource{
	//	Group:    "apps",
	//	Version:  "v1",
	//	Resource: "deployments",
	//}
	//deployment, err := client.Resource(gvr).Namespace("jenkins").Get(context.TODO(), "nginx", metav1.GetOptions{})
	//currentRevision := deployment.GetAnnotations()["deployment.kubernetes.io/revision"]
	//fmt.Println("current revision:", currentRevision)
	//historyRevision := "5"
	//patchData := fmt.Sprintf(`{"metadata":{"annotations":{"deployment.kubernetes.io/revision":"%s"}}}`, historyRevision)
	//if _, err = client.Resource(gvr).Namespace("jenkins").Patch(context.TODO(), "nginx", types.StrategicMergePatchType, []byte(patchData), metav1.PatchOptions{}); err != nil {
	//	log.Println("Patch", err)
	//	return
	//}
	//
	//fmt.Println("rollback done")
>>>>>>> 4622a0a (2025-1-15)
}
