package main

import (
	"flag"
	"fmt"
	"github.com/AliyunContainerService/reminder/internal"
	"github.com/AliyunContainerService/reminder/utils"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

var (
	configPath string
	kubeconfig string
)

func main() {
	stopper:=make(chan struct{})
	defer close(stopper)
	//
	flag.StringVar(&configPath, "path", "config/version.json", "pls input config path.")
	flag.StringVar(&kubeconfig,"kubeconfig","","Path to a kube config. Only required if out-of-cluster.")
	flag.Parse()

	cfg, err := utils.Load(configPath)
	if err != nil {
		panic(err)
	}

	clientset := internal.GetClientsetOrDie(kubeconfig)
	informerFactory := informers.NewSharedInformerFactory(clientset, 0)
	deployInformer:=informerFactory.Apps().V1().Deployments()
	defer runtime.HandleCrash()

	go informerFactory.Start(stopper)

	if !cache.WaitForCacheSync(stopper,deployInformer.Informer().HasSynced){
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	deploymentLister := deployInformer.Lister()
	deployList, err := deploymentLister.List(labels.Everything())
	if err != nil {
		klog.Error(err)
	}

	for _, deploy := range deployList {
		if deploy.APIVersion !=cfg.Deployment{
			fmt.Printf("The deployment %s's ApiVersion is not the %s,pls update quickly.\n",deploy.Name,cfg.Deployment)
		}
	}

	<-stopper
}
