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
	flag.StringVar(&JWTKey, "key", "testkey", "jwt key")
	flag.Parse()
	log.Printf("jwt key: %s \n", JWTKey)
}
