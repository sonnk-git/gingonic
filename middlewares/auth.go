package middlewares

import (
	"errors"
	"gingonic/db"
	"gingonic/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signed method received")
		}
		return []byte(os.Getenv("APP_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}
	return token, nil
}

func Build(user models.User) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	claims := jwt.MapClaims{
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(24 * 60 * time.Minute).Unix(),
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET_KEY")))

	return tokenString, err
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := Parse(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)

	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	exists := false
	err = db.Orm.Model(&models.User{}).Select("count(*) > 0").
		Where("id = ?", claims["id"]).
		Find(&exists).
		Error
	if err != nil || !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
	}

	c.Set("email", claims["email"])
	c.Set("username", claims["name"])
	c.Set("id", claims["id"])

	c.Next()
}

func JwtTokenCheckInGraphql(tokenString string) (models.User, error) {
	tokenString, err := ExtractBearerToken(tokenString)
	token, err := Parse(tokenString)
	user := models.User{}
	if err != nil {
		return user, err
	}
	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		return user, errors.New("unable to parse claims")
	}

	tx := db.Orm.First(&user, "id = ?", claims["id"])
	if tx.Error != nil {
		return user, errors.New("account not exists")
	}
	return user, err
}
