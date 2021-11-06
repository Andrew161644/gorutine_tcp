package main

import (
	"tcp/app/config"
	"tcp/app/connections_core"
)

func main() {
	var getConfig, _ = config.GetConfig()
	var server = connections_core.ConnectionsCore{}
	server.StartConnections(*getConfig)
	var ch = make(chan string)
	<-ch
}
