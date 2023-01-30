package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//将kubeconfig文件转化成rest.config类型的对象
	conf, err := clientcmd.BuildConfigFromFlags("", "D:\\config")
	if err != nil {
		panic(err)
	}
	//根据rest.config类型的对象，new一个clientset出来
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	//使用clientset获取pod列表
	podlist, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range podlist.Items {
		fmt.Println(pod.Name, pod.Namespace)
	}
}
