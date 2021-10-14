package pods

import (
	"context"
	"github.com/gin-gonic/gin"
	"hackube/env"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var config *rest.Config
var cs *kubernetes.Clientset

func init() {
	config, _ = rest.InClusterConfig()
	cs, _ = kubernetes.NewForConfig(config)
}

func RegisterPodRoutes(router *gin.Engine) {
	router.GET("/pods", Pods)
}

func Pods(c *gin.Context) {
	pl, _ := cs.CoreV1().Pods(string(env.ReadNamespace())).List(context.TODO(), metav1.ListOptions{})
	var names []string
	for _, item := range pl.Items {
		names = append(names, item.Name)
	}

	o, _ := json.Marshal(names)
	c.Writer.Write(o)
}
