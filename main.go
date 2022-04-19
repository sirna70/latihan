package main

import (
	"latihan/controllers"
	"latihan/helpers"

	"github.com/gin-gonic/gin"
)

var (
	__PORT__ = ":9090"
)

func main() {
	p := helpers.Post{}
	controllerPost := controllers.InPost{
		Post: &p,
	}

	router := gin.Default()

	router.GET("/posts", controllerPost.GetPost)
	router.GET("/posts/:id", controllerPost.GetPostById)
	router.POST("/posts", controllerPost.CreatePost)

	router.Run(__PORT__)
}
