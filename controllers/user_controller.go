package controllers

import (
	"belajar/config"
	"belajar/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

type AuthInputLogin struct{
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context){
	var input AuthInputRegister
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	
	// hash Passowrd
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if errHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "failed to hash password"})
		return
	}
	
	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
	}
	
	userCreated := config.DB.Create(&user).Error
	
	if userCreated != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "failed to create user"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message" : "user created successfully", "user": gin.H{
			"name": user.Name,
			"email": user.Email,
			"id": user.ID,
		}})
}