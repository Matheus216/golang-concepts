package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string
	Title  string
	Artist string
	Price  float64
}

var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var domain = "/albums/"

func main() {
	router := gin.Default()
	router.GET(domain, handleGetAll)
	router.POST(domain, handlePost)
	router.GET(domain+":id", handleGetById)

	router.Run("localhost:9999")
}

func handleGetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func handlePost(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func handleGetById(c *gin.Context) {
	id := c.Param("id")

	for _, item := range albums {
		if item.Id == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
