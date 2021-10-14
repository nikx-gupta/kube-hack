package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)

func AddEnvRoutes(router *gin.Engine) {
	router.GET("/env", PrintEnv)
	router.GET("/env/drives", Drives)
	router.GET("/env/mounts", Mounts)
	router.GET("/env/hosts", Hosts)
	router.GET("/env/hostname", HostName)
}

func HostName(c *gin.Context) {
	hst, _ := os.Hostname()
	c.Writer.WriteString(hst)
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

func ExecCommand(c *gin.Context, msg string, cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)
	data, err := cmd.Output()
	if err != nil {
		c.Error(err)
	}

	c.Writer.WriteString(msg)
	c.Writer.Write(data)
}
