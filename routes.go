package main

import "github.com/gin-gonic/gin"

func InitializeRoutes(router*gin.Engine){
	// user routes
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUser)
}