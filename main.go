package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	InitializeRoutes(router)
	router.Run(":8080")
}