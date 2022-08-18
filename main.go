package main

import (
	"restapi/controllers"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/api/albums", controllers.GetAlbums)
	router.GET("/api/albums/:id", controllers.GetAlbum)
	router.POST("/api/albums", controllers.CreateAlbum)
	router.PATCH("/api/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/api/albums/:id", controllers.DeleteAlbum)

	router.Run()
}
