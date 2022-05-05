package route

import (
	"github.com/bitbeliever/go-sso/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	e := gin.New()
	e.GET("/test", test())

	{
		group := e.Group("/user")
		group.POST("/login", login())
		group.POST("/signup", signup())
	}

	// auth group
	{
		group := e.Group("/auth")
		group.Use(middleware.Auth)
		//group.GET("/test", test())
		group.POST("/refresh_token", refreshToken())
	}

	return e
}
