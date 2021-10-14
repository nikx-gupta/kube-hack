package main

import (
	"github.com/gin-gonic/gin"
	"hackube/env"
	"hackube/pods"
	"hackube/services"
)

func main() {
	StartAPI()
}

func StartAPI() {
	router := gin.Default()

	router.GET("/ready", func(c *gin.Context) {
		c.Writer.WriteString("Container Ready To Serve")
	})
	router.GET("/healthy", func(c *gin.Context) {
		c.Writer.WriteString("ok")
	})

	env.RegisterRoutes(router)
	pods.RegisterPodRoutes(router)
	services.RegisterServiceRoutes(router)

	router.Run(":8080")
}
