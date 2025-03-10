package main

import (
	"log"

	"github.com/call-forwarding/config"
)

func main() {
	conf := config.GetConfig()
	log.Printf("config Server :: %+v", *conf.Server)
	log.Printf("config Db :: %+v", *conf.Db)
}
