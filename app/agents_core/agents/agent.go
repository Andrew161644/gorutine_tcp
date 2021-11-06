package agents

import (
	"log"
	"net"
)

type CommandEvent struct {
	Conn    *net.Conn
	Command map[string]interface{}
}

type IAgent interface {
	GetId() string
	GetStop() *chan struct{}
	PushCommand(commands CommandEvent)
	Unsubscribe()
	Handle(commands CommandEvent)
}

type Agent struct {
	Id   string
	Stop chan struct{}
}

func (r *Agent) Unsubscribe() {
	r.Stop <- struct{}{}
	log.Println("Stop func")
}

func (r *Agent) GetId() string {
	return r.Id
}

func (r *Agent) GetStop() *chan struct{} {
	return &r.Stop
}

func (*Agent) Handle(command CommandEvent) {
	log.Println("Base agent, trying execute: ", command.Command)
}

func (a *Agent) PushCommand(command CommandEvent) {

}

func (c *CommandEvent) GetCommandName() string {
	var keys []string
	for s := range c.Command {
		keys = append(keys, s)
	}
	return keys[0]
}
