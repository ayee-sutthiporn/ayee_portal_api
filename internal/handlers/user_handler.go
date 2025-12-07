package handlers

import (
	"net/http"
	"time"

	"ayee-portal-backend/internal/database"
	"ayee-portal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// SyncUser ensures the user exists in the local database
// SyncUser godoc
// @Summary      Sync user data
// @Description  Sync user data from frontend/auth provider to local DB
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/sync [post]
func SyncUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Upsert user
	var user models.User
	result := database.DB.Where("id = ?", input.ID).First(&user)
	if result.Error != nil {
		// Create
		input.CreatedAt = time.Now()
		input.UpdatedAt = time.Now()
		if err := database.DB.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Update
		user.Username = input.Username
		user.Email = input.Email
		user.Avatar = input.Avatar
		user.Role = input.Role
		user.UpdatedAt = time.Now()
		database.DB.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{"message": "User synced"})
}
