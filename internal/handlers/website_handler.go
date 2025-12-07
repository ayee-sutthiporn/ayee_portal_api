package handlers

import (
	"net/http"

	"ayee-portal-backend/internal/database"
	"ayee-portal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetWebsites godoc
// @Summary      Get all websites
// @Description  Get a list of all websites with their categories
// @Tags         websites
// @Produce      json
// @Success      200  {array}   models.Website
// @Failure      500  {object}  map[string]string
// @Router       /websites [get]
func GetWebsites(c *gin.Context) {
	var websites []models.Website
	// Preload Category
	if err := database.DB.Preload("Category").Find(&websites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, websites)
}

// CreateWebsite godoc
// @Summary      Create a new website
// @Description  Create a new website entry
// @Tags         websites
// @Accept       json
// @Produce      json
// @Param        website  body      models.Website  true  "Website Data"
// @Success      201      {object}  models.Website
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /websites [post]
func CreateWebsite(c *gin.Context) {
	var input models.Website
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// DeleteWebsite godoc
// @Summary      Delete a website
// @Description  Delete a website by ID
// @Tags         websites
// @Produce      json
// @Param        id   path      string  true  "Website ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /websites/{id} [delete]
func DeleteWebsite(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Website{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Website deleted"})
}
