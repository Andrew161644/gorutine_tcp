package commands

import "log"

type ICommand interface {
	Execute() interface{}
}

type CommandBase struct {
	A int
	B int
}

func (r CommandBase) Execute() interface{} {
	log.Println("Exec command", r)
	return nil
}

type CommandPlus struct {
	A int
	B int
}

func (r CommandPlus) Execute() interface{} {
	var res = r.A + r.B
	log.Println("Exec command plus", res)
	return res
}

type CommandMinus struct {
	A int
	B int
}

func (r CommandMinus) Execute() interface{} {
	var res = r.A - r.B
	log.Println("Exec command minus", res)
	return res
}
