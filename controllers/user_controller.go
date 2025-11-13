package controllers

import (
	"EFB-User-Backend/database"
	"EFB-User-Backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var input struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Role      string `json:"role"`
		AirlineID uint   `json:"airline_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hash),
		Role:      input.Role,
		AirlineID: input.AirlineID,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Airline").Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
