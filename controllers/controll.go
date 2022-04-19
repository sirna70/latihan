package controllers

import (
	"latihan/helpers"
	"latihan/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InPost struct {
	Post *helpers.Post
}

func (in *InPost) GetPost(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	posts, err := in.Post.GetPost("/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (in *InPost) GetPostById(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}
	id := c.Param("id")
	post, err := in.Post.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (in *InPost) CreatePost(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}
	var post helpers.Post
	c.BindJSON(&post)

	err := in.Post.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
