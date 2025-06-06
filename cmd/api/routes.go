package main

import (
	"microservice/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesCategory(router *gin.Engine) {

	categoryRoutes := router.Group("/categories")

	categoryRoutes.GET("/create", controllers.CategoryCreate)

}
