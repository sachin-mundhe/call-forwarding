package server

import (
	"fmt"

	"github.com/call-forwarding/harperDB/handlers"
)

func (e *echoServer) Start() {
	e.app.POST("/unconditional-call-forwardings", handlers.NewHarperHttpHandler(e.db).UnConditionalCF)
	e.app.POST("/call-forwardings", handlers.NewHarperHttpHandler(e.db).CallForwardings)

	serverUrl := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}
