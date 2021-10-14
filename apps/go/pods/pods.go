package pods

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackube/env"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

var config *rest.Config
var cs *kubernetes.Clientset

func init() {
	var err error
	config, err = rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	cs, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}

func RegisterPodRoutes(router *gin.Engine) {
	router.GET("/pods", Pods)
	router.GET("/pod/:podName/containers", Containers)
}

func Pods(c *gin.Context) {
	ns := strings.Replace(string(env.ReadNamespace()), "\n", "", 1)
	pl, err := cs.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var names []string
	for _, item := range pl.Items {
		names = append(names, item.Name)
	}

	o, _ := json.Marshal(names)
	c.Writer.Write(o)
}

func Containers(c *gin.Context) {
	pod_name := c.Param("podName")
	pl, err := cs.CoreV1().Pods(env.ReadNamespace()).Get(context.TODO(), pod_name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	var names []string
	for _, item := range pl.Spec.Containers {
		names = append(names, item.Name)
	}

	o, _ := json.Marshal(names)
	c.Writer.Write(o)
}
