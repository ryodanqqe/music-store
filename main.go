package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var storage = NewStorage()

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")
}

func getRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("albums/:id", deleteAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	return router
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := storage.ReadOne(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	err := storage.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, album{})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	storage.Create(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album
	c.BindJSON(&newAlbum)
	album, err := storage.Update(id, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
