package auth

import (
	"fmt"
	"gingonic/db"
	"gingonic/middlewares"
	"gingonic/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	user := &models.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}
	passInput := user.Password

	record := db.Orm.First(user, "email = ?", user.Email)
	var count int64
	record.Count(&count)
	if count == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "Account does not exist": user.Email})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passInput)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "Wrong password for": user.Email})
		return
	}

	// build token JWT
	token, err := middlewares.Build(*user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error when create token JWT"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
		"token":  token,
		"email":  user.Email,
	})
}

type EmailRequestBody struct {
	Email string `json:"email"`
}

func Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
	}

	fmt.Printf("%+v\n", user)
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	tx := db.Orm.Create(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passwordHashed),
	})
	if tx.Error != nil {
		_ = c.Error(tx.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": tx.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Register account successfully.",
	})
}

func Logout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}

func ForgotPassword(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}

func CheckIsLogin(c *gin.Context) {
	jwtToken, err := middlewares.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}

	token, err := middlewares.Parse(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)

	if !OK {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}

	exists := false
	err = db.Orm.Model(&models.User{}).Select("count(*) > 0").
		Where("id = ?", claims["id"]).
		Find(&exists).
		Error
	if err != nil || !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Valid token",
	})
}
