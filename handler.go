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

func createUser(c*gin.Context){
	var user User
	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}