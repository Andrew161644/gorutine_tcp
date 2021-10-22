package resolver

import (
	"log"
	"tcp/app/agents_core/agents"
)

type ServerResolver struct {
	Commands    chan agents.Commands
	Unsubscribe chan string
	AgentsChan  chan agents.IAgent
	Agents      map[string]agents.IAgent
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (r *ServerResolver) Close() {

}

func (r *ServerResolver) Start(rabbitMqConfig Config) {
	go r.GoBroadCastEvent()
}

func (r *ServerResolver) AddHandler(iSubscriber agents.IAgent) {
	r.AgentsChan <- iSubscriber
}

func (r *ServerResolver) GoBroadCastEvent() {
	for {
		select {
		case id := <-r.Unsubscribe:
			delete(r.Agents, id)
		case s := <-r.AgentsChan:
			r.Agents[s.GetId()] = s
		case commandMap := <-r.Commands:
			for id, s := range r.Agents {
				go func(id string, s agents.IAgent) {
					select {
					case <-*s.GetStop():
						r.Unsubscribe <- id
						return
					default:
						go s.Handle(commandMap)
					}
				}(id, s)
			}
		}
	}
}

func (r *ServerResolver) AddCommand(event agents.Commands) {
	r.Commands <- event
}
