package resolvers

import (
	"log"
	"reflect"
	"tcp/app/agents_core/agents"
)

type AgentResolver struct {
	Commands    chan agents.CommandEvent
	Unsubscribe chan string
	AgentsChan  chan agents.IAgent
	Agents      map[string]agents.IAgent
}

func (r *AgentResolver) Close() {

}

func (r *AgentResolver) Start(config Config) {
	go r.GoBroadCastEvent()
}

func (r *AgentResolver) AddAgent(iSubscriber agents.IAgent) {
	r.AgentsChan <- iSubscriber
}

func (r *AgentResolver) GoBroadCastEvent() {
	for {
		select {
		case id := <-r.Unsubscribe:
			delete(r.Agents, id)
		case s := <-r.AgentsChan:
			{
				r.Agents[s.GetId()] = s
				var agentName = reflect.TypeOf(s).Name()
				log.Println("Agent: ", agentName, " was added")
			}
		case commandEvent := <-r.Commands:
			{
				for id, s := range r.Agents {
					go func(id string, s agents.IAgent) {
						select {
						case <-*s.GetStop():
							r.Unsubscribe <- id
							return
						default:
							{
								var agentName = reflect.TypeOf(s).Name()
								if CanAgentExec(commandEvent, agentName) {
									go s.Handle(commandEvent)
								}
							}
						}
					}(id, s)
				}
			}
		}
	}
}

func (r *AgentResolver) AddCommand(event agents.CommandEvent) {
	r.Commands <- event
}

func CanAgentExec(commands agents.CommandEvent, agentName string) bool {

	var commandName string
	for s := range commands.Command {
		commandName = s
	}

	if agentName == "TestAgent" && commandName == "TestCommand" {
		return true
	}

	return false
}
