package main

import "github.com/gin-gonic/gin"

func main() {
	// creates a new instance of Gin with added middleware, logger, recovery middleware that catches an endpoint if it panics 
	r := gin.Default()

	// declare a route 
	r.GET("/ping", func(c *gin.Context) {
		// use the gin.Context to respond to the incoming request
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// run the router
	r.Run() 
}