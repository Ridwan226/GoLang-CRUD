package controllers

import (
	"belajar/config"
	"belajar/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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


func LoginUser(c *gin.Context){
	var input AuthInputLogin
	// Validation 
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
 var user models.User
 
 userData := config.DB.Where("email = ?", input.Email).First(&user).Error
	if userData != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email Not Found"})
		return
	}
	
	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	
	if errMatchPassword != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}
	
	// Buat Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub" : user.ID,
		"exp" : time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	
	tokenString, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if errToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "failed to create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"name": user.Name,
			"email": user.Email,
			"id": user.ID,
		}})
}