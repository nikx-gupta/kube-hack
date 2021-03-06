package env

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"os/exec"
	"strings"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/env", PrintEnv)
	router.GET("/env/drives", Drives)
	router.GET("/env/mounts", Mounts)
	router.GET("/env/hosts", Hosts)
	router.GET("/env/hostname", HostName)
	router.GET("/env/secrets", Secrets)
	router.GET("/env/ns", Namespace)
	router.GET("/env/token", Token)
	router.GET("/env/cert", Cert)
}

const Path_svc_acct = "/var/run/secrets/kubernetes.io/serviceaccount"
const ApiUrl = "https://localhost:16443"

var ns string
var KubeConfig *rest.Config
var KubeClientSet *kubernetes.Clientset

func init() {
	var err error
	KubeConfig, err = rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	KubeClientSet, err = kubernetes.NewForConfig(KubeConfig)
	if err != nil {
		panic(err)
	}
}

func HostName(c *gin.Context) {
	hst, _ := os.Hostname()
	c.Writer.WriteString(hst)
}

func Namespace(c *gin.Context) {
	c.Writer.WriteString(ReadNamespace())
}

func ReadNamespace() string {
	if ns == "" {
		data, _ := os.ReadFile(fmt.Sprintf("%s/namespace", Path_svc_acct))
		ns = strings.Replace(string(data), "\n", "", 1)
	}

	return ns
}

func ReadToken() []byte {
	data, _ := os.ReadFile(fmt.Sprintf("%s/token", Path_svc_acct))
	return data
}

func Cert(c *gin.Context) {
	data, _ := os.ReadFile(fmt.Sprintf("%s/ca.crt", Path_svc_acct))
	c.Writer.Write(data)
}

func Token(c *gin.Context) {
	c.Writer.Write(ReadToken())
}

func PrintEnv(c *gin.Context) {
	ExecCommand(c, "----- Printing Container Environment ----- \n", "printenv")
}

func Drives(c *gin.Context) {
	ExecCommand(c, "----- Printing Drives ----- \n", "df", "-h")
}

func Mounts(c *gin.Context) {
	ExecCommand(c, "----- Printing Mounts ----- \n", "findmnt")
}

func Hosts(c *gin.Context) {
	ExecCommand(c, "----- Printing Hosts ----- \n", "cat", "/etc/hosts")
}

func Secrets(c *gin.Context) {
	ExecCommand(c, "----- Printing Mounted Secrets ----- \n", "ls", "-alR", "/run/secrets")
}

func ExecCommand(c *gin.Context, msg string, cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)
	data, err := cmd.Output()
	if err != nil {
		c.Error(err)
	}

	c.Writer.WriteString(msg)
	c.Writer.Write(data)
}
