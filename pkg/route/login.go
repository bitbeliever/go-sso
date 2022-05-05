package route

import (
	"github.com/bitbeliever/go-sso/pkg/middleware"
	"github.com/bitbeliever/go-sso/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	errUnAuth       = response.Response{}
	errBadParameter = response.Response{
		Code:    400001,
		Message: "bad parameter",
		Data:    nil,
	}
)

func testToken() func(*gin.Context) {
	return func(c *gin.Context) {
		v, ok := c.Get("jwt_claim")
		if !ok {
			c.AbortWithStatus(500)
			return
		}
		u, ok := v.(middleware.User)
		if !ok {
			c.AbortWithStatus(500)
			return
		}

		c.JSON(200, response.Response{
			Code:    200,
			Message: "",
			Data:    u,
		})
	}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login() func(*gin.Context) {
	return func(c *gin.Context) {
		var u User
		if err := c.BindJSON(&u); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(400, errBadParameter)
			return
		}

		if u.Username == "xantares" && u.Password == "123" {

		}
	}
}

func refreshToken() func(*gin.Context) {
	return func(c *gin.Context) {
	}
}

// signup 注册
func signup() func(*gin.Context) {
	return func(c *gin.Context) {
	}
}

func test() func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := middleware.NewToken(middleware.User{
			UserID:   1,
			Username: "xantares",
			Password: "qwer",
			Email:    "aaa@ff.com",
			Avatar:   "noavatar",
		})
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		log.Println(token)
	}
}
