package main

import (
	"tcp/app/connections_core"
	"tcp/app/connections_core/config"
)

// требуется только ниже для обработки примера

//func main() {
//
//	fmt.Println("Launching server...")
//
//	// Устанавливаем прослушивание порта
//	ln, _ := net.Listen("tcp", ":8081")
//
//	// Открываем порт
//
//	go func() {
//		conn, _ := ln.Accept()
//		fmt.Println("Con1")
//		for {
//			// Будем прослушивать все сообщения разделенные \n
//			message, _ := bufio.NewReader(conn).ReadString('\n')
//			// Распечатываем полученое сообщение
//			fmt.Print("Message Received 1:", string(message))
//			// Процесс выборки для полученной строки
//			newmessage := strings.ToUpper(message)
//			// Отправить новую строку обратно клиенту
//			conn.Write([]byte(newmessage + "1\n"))
//		}
//		fmt.Println("Con2 closed")
//	}()
//
//	go func() {
//		conn, _ := ln.Accept()
//		fmt.Println("Con2")
//		for {
//			// Будем прослушивать все сообщения разделенные \n
//			message, _ := bufio.NewReader(conn).ReadString('\n')
//			// Распечатываем полученое сообщение
//			fmt.Print("Message Received 2:", string(message))
//			// Процесс выборки для полученной строки
//			newmessage := strings.ToUpper(message)
//			// Отправить новую строку обратно клиенту
//			conn.Write([]byte(newmessage + "2\n"))
//		}
//		fmt.Println("Con2 closed")
//	}()
//	var ch = make(chan string)
//	<-ch
//}

func main() {
	var server = connections_core.ConnectionsCore{}
	server.StartConnections(config.Config{
		Port:             ":8081",
		ConnectionsCount: "1",
	})
	var ch = make(chan string)
	<-ch
}
