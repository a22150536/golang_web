package alert

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	ctx       = context.Background()
)

type DeploymentInfo struct {
	Namespace string
	Name      string
	PodCount  int32
	Hpa       HpaInfo
}

type HpaInfo struct {
	Name   string
	MaxPod int32
	MinPod int32
}

func init() {
	kubeconfig := "/root/.kube/config"
	// kubeconfig := "D:/go/learn/src/ginmod1/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config", config)
	clientset, err = kubernetes.NewForConfig(config)
	fmt.Println("clientset", clientset)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *DeploymentInfo) GetDeployment() (err error) {
	fmt.Println("errrr4")
	deploymentsClient := clientset.AppsV1().Deployments(this.Namespace)
	fmt.Println("deploymentsClient", deploymentsClient)
	// v1, err := deploymentsClient.Get(ctx, this.Name, metav1.GetOptions{})
	fmt.Println("errrr15")
	v1, err := deploymentsClient.Get(ctx, this.Name, metav1.GetOptions{})
	fmt.Println("v1", v1)
	fmt.Println("errrr16")
	if err != nil {
		fmt.Println("setDeployment err = ", err)
		return
	}
	this.PodCount = v1.Status.Replicas
	fmt.Println("errrr10")
	return
}

func (this *DeploymentInfo) GetHPAResource() (err error) {
	fmt.Println("errrr5")
	hpaClient := clientset.AutoscalingV2().HorizontalPodAutoscalers(this.Namespace)
	info, err := hpaClient.Get(ctx, this.Hpa.Name, metav1.GetOptions{})
	if err != nil {
		fmt.Println("getHpa err = ", err)
		return
	}
	this.Hpa.MaxPod = info.Spec.MaxReplicas
	this.Hpa.MinPod = *info.Spec.MinReplicas
	return
}

func (this *DeploymentInfo) SetHPAResource() (err error) {
	fmt.Println("errrr6")
	hpaClient := clientset.AutoscalingV2().HorizontalPodAutoscalers(this.Namespace)
	info, err := hpaClient.Get(ctx, this.Hpa.Name, metav1.GetOptions{})
	this.PodCount = this.PodCount + 1
	fmt.Println("podcount = ", this.PodCount)
	info.Spec.MinReplicas = &this.PodCount
	fmt.Println("minreplicas = ", info.Spec.MinReplicas)
	result, err := hpaClient.Update(ctx, info, metav1.UpdateOptions{})
	fmt.Println("result = ", result.Status)
	if err != nil {
		fmt.Println("setHpa err = ", err)
	}
	return
}
