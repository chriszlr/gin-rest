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

func updateUser(c*gin.Context){
	id := c.Param("id")

	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users{
		if user.ID == id {
			// Update fields if they exist in the request body
			if updatedUser.Username != "" {
				users[i].Username = updatedUser.Username
			}
			if updatedUser.Email != "" {
				users[i].Email = updatedUser.Email
			}

			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c*gin.Context){
	id := c.Param("id")

	for i, user := range users{
		if user.ID == id{
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func getPosts(c*gin.Context){
	c.IndentedJSON(http.StatusOK, posts)
}

func createPost(c*gin.Context){
	var post Post

	if err := c.BindJSON(&post); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts = append(posts, post)
	c.JSON(http.StatusCreated, post)
}

func getPostById(c*gin.Context){
	id := c.Param("id")

	for _, post := range posts{
		if post.ID == id{
			c.JSON(http.StatusOK, post)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func updatePost(c*gin.Context){
	id := c.Param("id")

	var updatedPost Post

	if err := c.BindJSON(&updatedPost); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, post := range posts{
		if(post.ID == id){
			if updatedPost.Title != "" {
				posts[i].Title = updatedPost.Title
			}

			if(updatedPost.Content != ""){
				posts[i].Content = updatedPost.Content
			}

			c.JSON(http.StatusOK, posts[i])
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
	}
}

func deletePost(c*gin.Context){
	id := c.Param("id")

	for i, post := range posts{
		if post.ID == id{
			posts = append(posts[:i], posts[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}