package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samantamayuri/Crud-go/initializers"
	"github.com/samantamayuri/Crud-go/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully", "user": user})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set cookie
	c.SetCookie("auth", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "Logged in successfully"})

}

func ValidateUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(200, gin.H{"message": "User is valid", "user": user})
}
