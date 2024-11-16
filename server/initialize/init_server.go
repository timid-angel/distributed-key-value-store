package initialize

import (
	"bufio"
	"distributed-key-value-store/server/domain"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex
var DEADLINE = 100

func removeClient(conn net.Conn) {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	conn.Close()
}

func handleClient(conn net.Conn, controller domain.IController) {
	err := conn.SetDeadline(time.Now().Add(time.Second * time.Duration(DEADLINE)))
	if err != nil {
		log.Println("failed to set connection deadline:", err)
		removeClient(conn)
		return
	}

	reader := bufio.NewReader(conn)
	for {
		connErr := conn.SetDeadline(time.Now().Add(time.Second * time.Duration(DEADLINE)))
		if connErr != nil {
			log.Println("failed to set connection deadline:", connErr)
			removeClient(conn)
			return
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client disconnected:", err)
			conn.Close()
			return
		}

		response := controller.HandleRequest(message)
		response += "\n"
		conn.Write([]byte(response))
	}
}

func InitServer(PORT int, controller domain.IController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v\n", PORT, err.Error())
	}

	defer listener.Close()
	log.Println("Server listening on port:", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}

		mu.Lock()
		clients[conn] = true
		mu.Unlock()

		go handleClient(conn, controller)
	}
}
