package main

import (
	"log"

	"github.com/call-forwarding/config"
	"github.com/call-forwarding/database"
	"github.com/call-forwarding/server"
)

func main() {
	log.Println("Call forwarding demo")
	conf := config.GetConfig()
	db := database.NewHarperDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
