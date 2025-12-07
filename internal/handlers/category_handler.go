package handlers

import (
	"net/http"

	"ayee-portal-backend/internal/database"
	"ayee-portal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary      Get all categories
// @Description  Get a list of all categories ordered by 'order'
// @Tags         categories
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {object}  map[string]string
// @Router       /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Order("\"order\" asc").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category entry
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Category Data"
// @Success      201       {object}  models.Category
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /categories [post]
func CreateCategory(c *gin.Context) {
	var input models.Category
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
