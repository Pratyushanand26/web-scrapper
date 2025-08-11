package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratyushanand26/web-scrapper/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context, DB *gorm.DB) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user := db.User{Email: input.Email, Username: input.Username, Password: string(hashed)}

	result:=DB.Create(&user)

	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":result.Error.Error()})
	}
	c.JSON(http.StatusAccepted,gin.H{"message":"user created"})

}
