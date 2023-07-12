package main

type User struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type Post struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	AuthorID string `json:"authorId"`
}