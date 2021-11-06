package main

import (
	"encoding/json"
	"net"
	"reflect"
	"tcp/app/agents_core/commands"
)
import "fmt"
import "bufio"

func main() {
	var commandMap = make(map[string]commands.ICommand)
	var command = commands.TestCommand{Info: "Info"}
	commandMap[reflect.TypeOf(command).Name()] = command
	body, _ := json.Marshal(commandMap)
	fmt.Println(body)
	var str = string(body) + "\n"
	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	// Отправляем в socket

	fmt.Fprintf(conn, str)
	// Прослушиваем ответ
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

}
