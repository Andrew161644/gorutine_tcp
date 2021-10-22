package agents

import (
	"log"
)

type Commands struct {
	Command map[string]interface{}
}

type IAgent interface {
	GetId() string
	GetStop() *chan struct{}
	Unsubscribe()
	Handle(event Commands)
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

func (*Agent) Handle(event Commands) {
	log.Println("Base handler, trying execute: ", event.Command)
}

type RabbitHandlerPlus struct {
	Agent
}

func (*RabbitHandlerPlus) Handle(event Commands) {

}

type RabbitHandlerMinus struct {
	Agent
}

func (*RabbitHandlerMinus) Handle(event Commands) {

}
