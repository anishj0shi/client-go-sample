package pkg

import (
	"context"
	"log"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type NSClient interface {
	CreateNamespace(name string) error
}

var nsClient NSClient

func GetClient() NSClient {
	if nsClient != nil {
		return nsClient
	} else {
		nsClient = new()
		return nsClient
	}
}

func new() NSClient {
	client := &nsClientStruct{}
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	client.k8sClientSet = clientSet
	return client
}

type nsClientStruct struct {
	k8sClientSet *kubernetes.Clientset
}

func (n *nsClientStruct) CreateNamespace(name string) error {
	ns := &v1.Namespace{
		ObjectMeta: v12.ObjectMeta{
			Name: name,
		},
	}
	_, err := n.k8sClientSet.CoreV1().Namespaces().Create(context.TODO(), ns, v12.CreateOptions{})
	if err != nil {
		log.Printf( "Unable to Create Namespace. err: %+v", err)
		return err
	} else {
		log.Printf( "Namespace %s created", name)
		return nil
	}
}
