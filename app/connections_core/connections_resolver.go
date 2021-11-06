package connections_core

import (
	"bufio"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"log"
	"net"
	"strconv"
	"strings"
	"tcp/app/agents_core/agents"
	"tcp/app/agents_core/resolvers"
	"tcp/app/connections_core/config"
)

type IConn interface {
	StartConnections(count int)
}

type ConnectionsCore struct {
	Connections   map[int]*chan int
	AgentResolver resolvers.AgentResolver
}

func (con *ConnectionsCore) StartConnections(config config.Config) {

	con.AgentResolver = resolvers.AgentResolver{
		Commands:    make(chan agents.CommandEvent),
		Unsubscribe: make(chan string),
		AgentsChan:  make(chan agents.IAgent),
		Agents:      make(map[string]agents.IAgent),
	}

	StartAgentResolver(con)

	log.Println("Launching server...")
	var connCount, _ = strconv.Atoi(config.ConnectionsCount)
	con.Connections = make(map[int]*chan int, connCount)

	var port = config.Port

	ln, _ := net.Listen("tcp", port)

	go StartConnectionAsync(connCount, con, ln)

}

func StartAgentResolver(con *ConnectionsCore) {

	go con.AgentResolver.Start(resolvers.Config{})
	u, _ := uuid.NewV4()

	con.AgentResolver.AddAgent(agents.TestAgent{Agent: &agents.Agent{
		Id:   u.String(),
		Stop: make(chan struct{}),
	}})
}

func StartConnectionAsync(goroutineCountAvailable int, con *ConnectionsCore, ln net.Listener) {
	for {
		if goroutineCountAvailable > 0 {
			go func(connections map[int]*chan int, listener net.Listener, goroutineCountAvailable int) {
				log.Println("Start goroutine")
				var conn, _ = listener.Accept()
				var connNumber = len(connections)
				log.Println("Start Connection: ", connNumber)
				var ch = make(chan int)
				connections[connNumber] = &ch

				for {
					message, _ := bufio.NewReader(conn).ReadString('\n')
					if message != "" {
						log.Println(message)
						var commandsMap map[string]interface{}
						var err = json.Unmarshal([]byte(strings.TrimSpace(message)), &commandsMap)
						if err != nil {
							_, err := conn.Write([]byte("Error\n"))
							if err != nil {
								log.Println(err)
								break
							}
							continue
						}
						var commandEvent = agents.CommandEvent{
							Conn:    &conn,
							Command: commandsMap,
						}

						if commandEvent.GetCommandName() == "CommandStop" {
							break
						}

						// Send commandEvent to AgentResolver
						con.AgentResolver.AddCommand(commandEvent)
					}
				}

				conn.Close()
				delete(connections, connNumber)
				log.Printf("Conn %d closed", connNumber)
				connNumber--
				log.Printf("Conn at server %d ", len(connections))
				goroutineCountAvailable++
				log.Printf("Goroutine available %d", goroutineCountAvailable)

			}(con.Connections, ln, goroutineCountAvailable)
			goroutineCountAvailable--
			log.Printf("Goroutine available %d", goroutineCountAvailable)
		}
	}
}
