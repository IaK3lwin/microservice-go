package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	RoutesCategory(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
