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

func getUserByID(c*gin.Context){
	id := c.Param("id")

	for _, user := range users{
		if user.ID == id{
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}