package main

import "net"
import "fmt"
import "bufio"
import "strings" // требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":8081")

	// Открываем порт
	conn, _ := ln.Accept()
	conn2, _ := ln.Accept()

	go func(conn net.Conn) {
		fmt.Println("Con1")
		for {
			// Будем прослушивать все сообщения разделенные \n
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// Распечатываем полученое сообщение
			fmt.Print("Message Received 1:", string(message))
			// Процесс выборки для полученной строки
			newmessage := strings.ToUpper(message)
			// Отправить новую строку обратно клиенту
			conn.Write([]byte(newmessage + "1\n"))
		}
		fmt.Println("Con2 closed")
	}(conn)

	go func(conn net.Conn) {
		fmt.Println("Con1")
		for {
			// Будем прослушивать все сообщения разделенные \n
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// Распечатываем полученое сообщение
			fmt.Print("Message Received 1:", string(message))
			// Процесс выборки для полученной строки
			newmessage := strings.ToUpper(message)
			// Отправить новую строку обратно клиенту
			conn.Write([]byte(newmessage + "1\n"))
		}
		fmt.Println("Con2 closed")
	}(conn2)
}
