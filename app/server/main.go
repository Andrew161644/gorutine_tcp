package main

import (
	"tcp/app/connections_core"
	"tcp/app/connections_core/config"
)

func main() {
	var server = connections_core.ConnectionsCore{}
	server.StartConnections(config.Config{
		Port:             ":8081",
		ConnectionsCount: "4",
	})
	var ch = make(chan string)
	<-ch
}
