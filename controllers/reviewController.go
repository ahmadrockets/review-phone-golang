package controllers

import (
	"net/http"

	"final-project-sanber/models"
	"final-project-sanber/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	PhoneID uint   `json:"phoneID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Picture string `json:"picture"`
	UserID  uint   `json:"userID"`
}

// GetAllReview godoc
// @Summary Get all Review.
// @Description Get a list of Review.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var reviews []models.Review
	db.Find(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// GetAllReviewByPhoneID godoc
// @Summary Get all Review By Phone ID.
// @Description Get a list of By Phone ID.
// @Tags Review
// @Param id path string true "Phone id"
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews/byphone/{id} [get]
func GetAllReviewByPhoneID(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var reviews []models.Review
	db.Where("phone_id = ?", c.Param("id")).Find(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body reviewInput true "the body to create a new Review"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate url picture
	urlPicture := ""
	if input.Picture != "" {
		urlval := utils.UrlValidation(input.Picture)
		if urlval == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "picture is not valid url"})
			return
		}
		urlPicture = input.Picture
	}

	// Create review
	review := models.Review{
		PhoneID: input.PhoneID,
		Title:   input.Title,
		Content: input.Content,
		Picture: urlPicture,
		UserID:  input.UserID,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
