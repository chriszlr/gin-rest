package main

import "github.com/gin-gonic/gin"

func InitializeRoutes(router*gin.Engine){
	// user routes
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)


	// post routes
	router.GET("/posts", getPosts)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", getPostById)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
}