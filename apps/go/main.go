package main

import "github.com/gin-gonic/gin"

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

	AddEnvRoutes(router)

	router.Run(":8080")
}
