package tests

import (
	"fmt"
	"log"
	"reflect"
	"tcp/app/agents_core/agents"
	"tcp/app/agents_core/commands"
	"tcp/app/agents_core/resolvers"
	"testing"
)

func TestCanMapCommand(t *testing.T) {
	var agent agents.IAgent = agents.TestAgent{}
	var agentName = reflect.TypeOf(agent).Name()
	log.Println(agentName)
	var command = commands.TestCommand{}
	var name = reflect.TypeOf(command).Name()
	log.Println(name)
	var commandMap = agents.CommandEvent{Command: map[string]interface{}{
		name: command,
	}}
	var res = resolvers.CanAgentExec(commandMap, agentName)
	if !res {
		t.Fatal("Error")
	}
}

func TestCanMapCommandFaild(t *testing.T) {
	var agent agents.IAgent = agents.TestAgent{}
	var agentName = reflect.TypeOf(agent).Name()
	fmt.Println(agentName)
	var command = commands.TestCommand{}
	var name = reflect.TypeOf(command).Name()
	log.Println(name)
	var commandMap = agents.CommandEvent{Command: map[string]interface{}{
		name: command,
	}}
	var res = resolvers.CanAgentExec(commandMap, agentName)
	if !res {
		t.Fatal("Error")
	}
}

func TestCanMapCommandReflect(t *testing.T) {
	var commandMap = agents.CommandEvent{
		Command: map[string]interface{}{
			"CommandStop": commands.CommandStop{},
		}}
	fmt.Println(reflect.TypeOf(commandMap).Name())
}
