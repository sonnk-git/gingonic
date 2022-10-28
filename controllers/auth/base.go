package auth

import (
	"encoding/json"
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
	//x, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("%s", string(x))
	//c.JSON(http.StatusOK, c)

	var requestBody EmailRequestBody

	if err := json.Compact().Unmarshal(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		println("loi roi" + err.Error())
	}
	fmt.Println("requestBody.Email")
	fmt.Println(requestBody.Email)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
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
}

func Logout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}

func ForgotPassword(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}
