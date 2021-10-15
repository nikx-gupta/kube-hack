package services

import (
	"github.com/gin-gonic/gin"
)

func RegisterServiceRoutes(router *gin.Engine) {
	//router.GET("/pods", Pods)
	//router.GET("/pod/:podName/containers", Containers)

}
//
//func PortForward(c *gin.Context){
//	svcName := c.Param("svcName")
//	svc, err := env.KubeClientSet.CoreV1().Services(env.ReadNamespace()).
//		Get(context.TODO(),svcName,metav1.GetOptions{})
//
//		//List(context.TODO(), metav1.ListOptions{})
//
//}
