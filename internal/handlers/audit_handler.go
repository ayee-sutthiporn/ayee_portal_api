package handlers

import (
	"net/http"
	"time"

	"ayee-portal-backend/internal/database"
	"ayee-portal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetAuditLogs godoc
// @Summary      Get audit logs
// @Description  Get a list of all audit logs ordered by timestamp
// @Tags         audit-logs
// @Produce      json
// @Success      200  {array}   models.AuditLog
// @Failure      500  {object}  map[string]string
// @Router       /audit-logs [get]
func GetAuditLogs(c *gin.Context) {
	var logs []models.AuditLog
	if err := database.DB.Order("timestamp desc").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// CreateAuditLog godoc
// @Summary      Create an audit log
// @Description  Create a new audit log entry
// @Tags         audit-logs
// @Accept       json
// @Produce      json
// @Param        log  body      models.AuditLog  true  "Audit Log Data"
// @Success      201  {object}  models.AuditLog
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /audit-logs [post]
func CreateAuditLog(c *gin.Context) {
	var input models.AuditLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Force timestamp
	input.Timestamp = time.Now()

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}
