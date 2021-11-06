package agents

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"tcp/app/agents_core/commands"
)

type TestAgent struct {
	*Agent
}

func (TestAgent) Handle(c CommandEvent) {
	log.Println("Test agent, trying execute: ", c.Command)
	var command commands.TestCommand
	for _, value := range c.Command {
		err := mapstructure.Decode(value, &command)
		if err != nil {
			log.Println(err)
			break
		}
	}
	var res = fmt.Sprintf("%v", command.Execute())

	log.Println(res)
	_, err := (*c.Conn).Write([]byte(res + "\n"))
	if err != nil {
		log.Println(err)
	}
}
