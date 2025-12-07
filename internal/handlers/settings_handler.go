package handlers

import (
	"net/http"

	"ayee-portal-backend/internal/database"
	"ayee-portal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetSystemSettings godoc
// @Summary      Get system settings
// @Description  Get the global system settings
// @Tags         settings
// @Produce      json
// @Success      200  {object}  models.SystemSettings
// @Failure      500  {object}  map[string]string
// @Router       /settings [get]
func GetSystemSettings(c *gin.Context) {
	var settings models.SystemSettings
	// Check if exists, if not create default
	if err := database.DB.First(&settings, 1).Error; err != nil {
		// Create default
		settings = models.SystemSettings{ID: 1}
		database.DB.Create(&settings)
	}
	c.JSON(http.StatusOK, settings)
}

// UpdateSystemSettings godoc
// @Summary      Update system settings
// @Description  Update the global system settings
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        settings  body      models.SystemSettings  true  "Settings Data"
// @Success      200       {object}  models.SystemSettings
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /settings [put]
func UpdateSystemSettings(c *gin.Context) {
	var input models.SystemSettings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Always update ID 1
	input.ID = 1

	if err := database.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}
