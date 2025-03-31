package main

import (
	"github.com/call-forwarding/config"
	"github.com/call-forwarding/database"
	"github.com/call-forwarding/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewHarperDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
