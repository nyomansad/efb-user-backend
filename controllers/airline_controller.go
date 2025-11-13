package controllers

import (
	"EFB-User-Backend/database"
	"EFB-User-Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAirline(c *gin.Context) {
	var airline models.Airline
	if err := c.ShouldBindJSON(&airline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&airline); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Airline created successfully", "data": airline})
}

func GetAirlines(c *gin.Context) {
	var airlines []models.Airline
	database.DB.Find(&airlines)
	c.JSON(http.StatusOK, gin.H{"airlines": airlines})
}

func DeleteAirline(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Airline{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Airline deleted"})
}
