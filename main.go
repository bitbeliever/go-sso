package main

import (
	"github.com/bitbeliever/go-sso/pkg/route"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if os.Getenv("GIN_RELEASE") != "" {
		gin.SetMode(gin.ReleaseMode)
	}

	g := route.Route()

	if err := g.Run(":8080"); err != nil {
		log.Println(err)
	}
}
