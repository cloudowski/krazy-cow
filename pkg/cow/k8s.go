package cow

import (
	"fmt"
	"log"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	pods      []*corev1.Pod
	endpoints []*corev1.Endpoints
)

const cowLabelSelector = "app=cow,herd=blue"

func getK8sClient() error {
	// // creates the in-cluster config
	var config *rest.Config
	var err error
	config, err = rest.InClusterConfig()
	if err != nil {
		logger.Error(err.Error())

		// try out-of-cluster
		// kubeconfig := "/Users/tomasz/projects/lvo/minikube-aws/kubeconfig.aws"
		kubeconfig := "/Users/tomasz/.kube/config"

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Panic(err.Error())
			return err
		}
	}

	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func GetK8sVersion() string {
	if err := getK8sClient(); err != nil {
		logger.Warningf("FAILED to connect to Kubernetes: %v", err)
		return ""
	}

	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		logger.Error(err)
	}
	return fmt.Sprintf("%s", version)
}

func GetCows() {

	if err := getK8sClient(); err != nil {
		logger.Errorf("FAILED to connect to Kubernetes: %v", err)
		return
	}

	log.Printf("KUBERNETES VERSION: %s", GetK8sVersion())

	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{LabelSelector: cowLabelSelector})
		if err != nil {
			logger.Errorf(err.Error())
		}

		logger.Infof("THERE are %d Pods:\n", len(pods.Items))
		for _, v := range pods.Items {
			logger.Infof("%v IP: %v PHASE: %v", v.Name, v.Status.PodIP, v.Status.Phase)

		}

		// svcs, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
		// if err != nil {
		// 	log.Print(err.Error())
		// }
		// log.Printf("Services:\n")
		// for _, v := range svcs.Items {
		// 	log.Printf("%v@%v %v: %v ", v.Name, v.Namespace, v.Spec.Type, v.Spec.ClusterIP)
		// }

		eps, err := clientset.CoreV1().Endpoints("").List(metav1.ListOptions{LabelSelector: cowLabelSelector})
		if err != nil {
			logger.Error(err.Error())
		}

		logger.Infof("Endpoints:\n")
		for _, v := range eps.Items {
			for _, e := range v.Subsets {
				for _, a := range e.Addresses {
					logger.Infof("Addresses:%v %v", a.IP, a.Hostname)
				}
			}
		}

		time.Sleep(time.Second * 30)
	}
}
