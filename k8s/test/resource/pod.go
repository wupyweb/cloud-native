package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	var kubeconfig string
	//if home := homedir.HomeDir(); home != "" {
	//	// 默认寻找~/.kube/config配置
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	kubeconfig = "client-go.kubeconfig"
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		//pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		//if err != nil {
		//	panic(err.Error())
		//}
		//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		//namespace := "default"
		//pod := "example-xxxxx"
		//_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		//if errors.IsNotFound(err) {
		//	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		//	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
		//		pod, namespace, statusError.ErrStatus.Message)
		//} else if err != nil {
		//	panic(err.Error())
		//} else {
		//	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		//}

		namespace := "default"

		// Examples for listing pods in default namespace
		//pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(),metav1.ListOptions{})
		//if err != nil {
		//	panic(err)
		//}
		//for _, pod := range pods.Items {
		//	fmt.Printf("%s\t%s\n",pod.Name, pod.Status.PodIP)
		//}

		// Examples for listing deployments in default namespace
		deploys, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, deploy := range deploys.Items {
			fmt.Printf("%s\t%d/%d\n", deploy.Name, deploy.Status.ReadyReplicas, deploy.Status.Replicas)
		}

		time.Sleep(10 * time.Second)
	}
}
