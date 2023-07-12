package main

import "github.com/gin-gonic/gin"

func InitializeRoutes(router*gin.Engine){
	// user routes
	router.GET("/users", getUsers)
}