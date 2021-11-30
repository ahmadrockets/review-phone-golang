package controllers

import (
	"net/http"

	"final-project-sanber/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roleInput struct {
	Name     string `json:"name"`
	CanWrite *bool  `json:"can_write"`
}

// GetAllRole godoc
// @Summary Get all Role.
// @Description Get a list of Role.
// @Tags Role
// @Produce json
// @Success 200 {object} []models.Role
// @Router /roles [get]
func GetAllRole(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var roles []models.Role
	db.Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// CreateRole godoc
// @Summary Create New Role.
// @Description Creating a new Role.
// @Tags Role
// @Param Body body roleInput true "the body to create a new Role"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Role
// @Router /roles [post]
func CreateRole(c *gin.Context) {
	// Validate input
	var input roleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create role
	role := models.Role{Name: input.Name, CanWrite: input.CanWrite}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&role)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// UpdateRole godoc
// @Summary Update Role.
// @Description Update Role by id.
// @Tags Role
// @Produce json
// @Param id path string true "Role id"
// @Param Body body roleInput true "the body to update role"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Role
// @Router /roles/{id} [patch]
func UpdateRole(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var role models.Role
	if err := db.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input roleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Role
	updatedInput.Name = input.Name
	updatedInput.CanWrite = input.CanWrite

	db.Model(&role).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// DeleteRole godoc
// @Summary Delete one Role.
// @Description Delete a Role by id.
// @Tags Role
// @Produce json
// @Param id path string true "Role id"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /roles/{id} [delete]
func DeleteRole(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var role models.Role
	if err := db.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
