package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	PORT := 8080
	conn, err := net.Dial("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatalln("failed to connect to server:", err)
	}

	defer conn.Close()

	go receiveMessage(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\033[92m> \033[0m")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}

func receiveMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil && strings.Contains(err.Error(), "timeout") {
			log.Fatalln("timeout", err)
			return
		}

		if err != nil {
			log.Fatalln("failed to send message:", err)
			return
		}

		fmt.Println("\033[92m\n\n\tMessage from server: \033[0m", message)
		fmt.Print("Send message to server: ")
	}
}
