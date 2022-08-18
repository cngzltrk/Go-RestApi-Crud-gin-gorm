package controllers

import (
	"net/http"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

type CreateAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type UpdateAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//GET /albums
func GetAlbums(c *gin.Context) {

	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// GET /albums/:id
func GetAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// POST /albums
func CreateAlbum(c *gin.Context) {
	// Validate input
	var input CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	album := models.Album{Title: input.Title, Artist: input.Artist, Price: input.Price}
	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// PATCH /albums/:id
func UpdateAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// DELETE /albums/:id
func DeleteAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
