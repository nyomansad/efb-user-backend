package controllers

import (
	"EFB-User-Backend/database"
	"EFB-User-Backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var input struct {
		UserID     uint   `json:"user_id"`
		DeviceInfo string `json:"device_info"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := models.UserSession{
		UserID:     input.UserID,
		LoginTime:  time.Now(),
		DeviceInfo: input.DeviceInfo,
	}

	database.DB.Create(&session)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login time recorded",
		"session": session,
	})
}

func UserLogout(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var session models.UserSession
	result := database.DB.
		Where("user_id = ? AND logout_time IS NULL", input.UserID).
		Order("id desc").
		First(&session)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No active login session found",
		})
		return
	}

	now := time.Now()
	session.LogoutTime = &now

	database.DB.Save(&session)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout time recorded",
		"session": session,
	})
}
