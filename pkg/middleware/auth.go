package middleware

import (
	"errors"
	"fmt"
	"github.com/bitbeliever/go-sso/config"
	"github.com/bitbeliever/go-sso/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type UserClaim struct {
	jwt.StandardClaims
	User
}

type User struct {
	UserID   int64
	Username string
	Password string
	Email    string
	Avatar   string
}

var (
	errUnAuth = response.Response{
		Code:    401,
		Message: "un auth",
		Data:    nil,
	}

	errTokenInvalid = response.Response{
		Code:    401,
		Message: "token is invalid",
		Data:    nil,
	}
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errUnAuth)
		return
	}

	user, err := jwtParse(token)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errTokenInvalid)
		return
	}

	log.Println(user)
	c.Set("jwt_claim", user)
	c.Next()
}

func jwtParse(tokenStr string) (User, error) {
	var userClaim UserClaim
	t, err := jwt.ParseWithClaims(tokenStr, &userClaim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTKey), nil
	})
	if err != nil {
		return User{}, err
	}

	if !t.Valid {
		return User{}, errors.New("jwt invalid")
	}

	return userClaim.User, nil
}

func NewToken(u User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{User: u})
	return token.SignedString([]byte(config.JWTKey))
}
