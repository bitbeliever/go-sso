package config

import (
	"flag"
	"log"
)

var (
	JWTKey string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	JWTKey = *flag.String("key", "test", "jwt key")

	flag.Parse()

	log.Printf("jwt key: %s \n", JWTKey)
}
