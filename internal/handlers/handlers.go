package handlers

import (
	"net/http"
	"strconv"

	"github.com/andkolbe/chirper-gin/internal/models"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	chirps, _ := models.ShowAllChirps()
	
	// call render function with the same name of the template to render
	render(c, gin.H{"title": "Home Page", "payload": chirps}, "home.html")
}

func GetChirp(c *gin.Context) {
	// check if the chirp ID is valid
	if chirpID, err := strconv.Atoi(c.Param("chirp_id")); err == nil {
		// check if the chirp exists
		if chirp, err := models.GetChirpByID(chirpID); err == nil {
			// Call the render function with the title, chirp and the name of the template
			render(c, gin.H{"title": chirp.Content, "payload": chirp}, "chirp.html")
		} else {
			// if the chirp is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid chirp ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}