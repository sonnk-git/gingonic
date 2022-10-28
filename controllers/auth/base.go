package auth

import (
	"fmt"
	"gingonic/db"
	"gingonic/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
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
		}

	c.IndentedJSON(http.StatusOK, struct{}{})
}

func Logout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}

func ForgotPassword(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}
