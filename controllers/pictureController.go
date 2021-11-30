package controllers

import (
	"net/http"

	"final-project-sanber/models"
	"final-project-sanber/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pictureInput struct {
	PhoneID    uint   `json:"phoneID"`
	Title      string `json:"title"`
	UrlPicture string `json:"urlPicture"`
}

// GetAllPicture godoc
// @Summary Get all Picture.
// @Description Get a list of Picture.
// @Tags Picture
// @Produce json
// @Success 200 {object} []models.Picture
// @Router /pictures [get]
func GetAllPicture(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var pictures []models.Picture
	db.Find(&pictures)

	c.JSON(http.StatusOK, gin.H{"data": pictures})
}

// GetAllPictureByPhoneID godoc
// @Summary Get all Picture By Phone ID.
// @Description Get a list of By Phone ID.
// @Tags Picture
// @Param id path string true "Phone id"
// @Produce json
// @Success 200 {object} []models.Picture
// @Router /pictures/byphone/{id} [get]
func GetAllPictureByPhoneID(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var pictures []models.Picture
	db.Where("phone_id = ?", c.Param("id")).Find(&pictures)

	c.JSON(http.StatusOK, gin.H{"data": pictures})
}

// CreatePicture godoc
// @Summary Create New Picture.
// @Description Creating a new Picture.
// @Tags Picture
// @Param Body body pictureInput true "the body to create a new Picture"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Picture
// @Router /pictures [post]
func CreatePicture(c *gin.Context) {
	// Validate input
	var input pictureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate url picture
	urlPicture := ""
	if input.UrlPicture != "" {
		urlval := utils.UrlValidation(input.UrlPicture)
		if urlval == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "picture is not valid url"})
			return
		}
		urlPicture = input.UrlPicture
	}

	// Create picture
	picture := models.Picture{PhoneID: input.PhoneID, Title: input.Title, UrlPicture: urlPicture}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&picture)

	c.JSON(http.StatusOK, gin.H{"data": picture})
}

// UpdatePicture godoc
// @Summary Update Picture.
// @Description Update Picture by id.
// @Tags Picture
// @Produce json
// @Param id path string true "Picture id"
// @Param Body body pictureInput true "the body to update picture"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Picture
// @Router /pictures/{id} [patch]
func UpdatePicture(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var picture models.Picture
	if err := db.Where("id = ?", c.Param("id")).First(&picture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input pictureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate url picture
	urlPicture := ""
	if input.UrlPicture != "" {
		urlval := utils.UrlValidation(input.UrlPicture)
		if urlval == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "picture is not valid url"})
			return
		}
		urlPicture = input.UrlPicture
	}

	var updatedInput models.Picture
	updatedInput.PhoneID = input.PhoneID
	updatedInput.Title = input.Title
	updatedInput.UrlPicture = urlPicture

	db.Model(&picture).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": picture})
}

// DeletePicture godoc
// @Summary Delete one Picture.
// @Description Delete a Picture by id.
// @Tags Picture
// @Produce json
// @Param id path string true "Picture id"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /pictures/{id} [delete]
func DeletePicture(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var picture models.Picture
	if err := db.Where("id = ?", c.Param("id")).First(&picture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&picture)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
