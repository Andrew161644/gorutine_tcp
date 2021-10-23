package connections_core

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"tcp/app/connections_core/config"
)

type IConn interface {
	StartConnections(count int)
}

type ConnectionsCore struct {
	Connections map[int]*chan int
}

func (con *ConnectionsCore) StartConnections(config config.Config) {
	fmt.Println("Launching server...")
	var connCount, _ = strconv.Atoi(config.ConnectionsCount)
	con.Connections = make(map[int]*chan int, connCount)

	var port = config.Port
	ln, _ := net.Listen("tcp", port)

	for i := 0; i < connCount; i++ {
		go func(connections map[int]*chan int, listener net.Listener) {
			var conn, _ = listener.Accept()
			var connNumber = len(connections)
			fmt.Println("Start Connection: ", connNumber)
			var ch = make(chan int)
			connections[connNumber] = &ch

			for {
				message, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Println("Message Received:", strings.TrimSpace(message))
				if strings.TrimSpace(message) == "stop" {
					break
				}
				newmessage := strings.ToUpper(message)
				conn.Write([]byte(newmessage + "1\n"))
			}
			conn.Close()
			fmt.Printf("Conn %d closed", connNumber)
		}(con.Connections, ln)
	}
}
