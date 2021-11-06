package main

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"reflect"
	"tcp/app/agents_core/commands"
)
import "fmt"

func main() {
	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		reader.ReadString('\n')
		Publish(conn, commands.TestCommand{Info: "Info"})
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func Publish(conn net.Conn, command commands.ICommand) {
	var commandMap = make(map[string]commands.ICommand)
	commandMap[reflect.TypeOf(command).Name()] = command
	body, _ := json.Marshal(commandMap)
	fmt.Fprintf(conn, string(body)+"\n")
}
