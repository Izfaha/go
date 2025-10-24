package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data
type album struct {
	ID		string  `json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},	
}

// get album endpoint as JSON
func getAlbums(ci *gin.Context){		// gin.Context = carry request detail
	ci.JSON(http.StatusOK, albums)	// Context.IndentedISON = serialize the struct into JSON and add it to the response.
}

// post Album
func postAlbums(ce *gin.Context) {
	var newAlbum album

	/* 
	*	call bindJSON to bind the receive JSON to
	*	newAlbum
	*/ 
	if err := ce.BindJSON(&newAlbum); err != nil {
		return
	}

	// add the new album to the slice
	albums = append(albums, newAlbum)
	ce.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
* getAlbumByID locates the album whose ID value matches the id
* parameter sent by the client, then returns that album as a response
*/ 
func getAlbumByID(cg *gin.Context) {
	id := cg.Param("id")

	/*
	* loop over the list of albums, looking for
	* an album whose ID value matches the parameter.
	*/ 
	for _, a := range albums {
		if a.ID == id {
			cg.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	cg.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// main function
func main() {
	router := gin.Default()

	// api endpoint for album
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// api endpoint for add new album
	router.POST("/albums", postAlbums)

	// Run = attach the router to an http.Server and start the server
	router.Run(":8345")
}