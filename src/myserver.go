/*
Theodore Ikehara
tcp echo server
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter the port to listen on: ")
	reader := bufio.NewReader(os.Stdin)
	portString, _ := reader.ReadString('\n')
	port, _ := strconv.Atoi(portString[:len(portString)-1])

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accepted connection from", conn.RemoteAddr())

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Print("Message Received:", string(message))
		newmessage := message
		conn.Write([]byte(newmessage + "\n"))
	}
}
