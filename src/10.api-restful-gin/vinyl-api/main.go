package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbum(ctx *gin.Context) {
	var album album

	if err := ctx.BindJSON(&album); err != nil {
		return
	}

	id := ctx.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums[i] = album
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
}

func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
}

func deleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Álbum eliminado con éxito"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
}

func main() {
	fmt.Println("Bienvenido a Vinyl-API")

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", putAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}
