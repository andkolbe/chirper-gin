package main

import "github.com/gin-gonic/gin"

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{"title": "Home Page"})
	})
}