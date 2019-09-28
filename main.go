package main

import (
	"flag"
	"fmt"
	clientset "github.com/liuyehcf/demo-controller/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	//klog.InitFlags(nil)
	//flag.Parse()
	//
	//// set up signals so we handle the first shutdown signal gracefully
	//stopCh := signals.SetupSignalHandler()
	//
	//cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	//if err != nil {
	//	klog.Fatalf("Error building kubeconfig: %s", err.Error())
	//}
	//
	//kubeClient, err := kubernetes.NewForConfig(cfg)
	//if err != nil {
	//	klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	//}
	//
	//exampleClient, err := clientset.NewForConfig(cfg)
	//if err != nil {
	//	klog.Fatalf("Error building example clientset: %s", err.Error())
	//}
	//
	//kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	//exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)
	//
	//controller := controllers.NewController(kubeClient, exampleClient,
	//	exampleInformerFactory.Democontroller().V1().MyConfigs())
	//
	//// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	//// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	//kubeInformerFactory.Start(stopCh)
	//exampleInformerFactory.Start(stopCh)
	//
	//if err = controller.Run(2, stopCh); err != nil {
	//	klog.Fatalf("Error running controller: %s", err.Error())
	//}

	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
		return
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
		return
	}

	myConfigList, err := exampleClient.DemocontrollerV1().MyConfigs("").List(metav1.ListOptions{})
	if err != nil {
		klog.Fatalf("Error listing myconfigs: %s", err.Error())
		return
	}

	fmt.Println(myConfigList)
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
