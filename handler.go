package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// user handlers
var users []User
var posts []Post

func getUsers(c*gin.Context){
	c.IndentedJSON(http.StatusOK, users)
}