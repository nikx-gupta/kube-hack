package pods

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackube/env"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"strings"
)

var currentPod *v1.Pod
var currentPodName string

func init() {
	currentPodName = os.Getenv("HOSTNAME")
	var err error
	currentPod, err = env.KubeClientSet.CoreV1().Pods(env.ReadNamespace()).
		Get(context.TODO(), currentPodName, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}
}

func RegisterPodRoutes(router *gin.Engine) {
	router.GET("/pods", Pods)
	router.GET("/pod/:podName/containers", Containers)

	router.GET("/pod/current", Current)
}

func Current(c *gin.Context) {
	data, _ := json.Marshal(currentPod)
	c.Writer.Write(data)
}

func Pods(c *gin.Context) {
	ns := strings.Replace(string(env.ReadNamespace()), "\n", "", 1)
	pl, err := env.KubeClientSet.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
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
	pl, err := env.KubeClientSet.CoreV1().Pods(env.ReadNamespace()).Get(context.TODO(), pod_name, metav1.GetOptions{})
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
