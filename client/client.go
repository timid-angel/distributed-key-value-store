package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var PORT = 8080
var conn net.Conn

func getConnection(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("failed to connect to server:", err)
	}

	return conn
}

func handleConnectionError(conn net.Conn, connectionReader *bufio.Reader, err error) (net.Conn, bool) {
	if err == nil {
		return conn, false
	}

	if !strings.Contains(err.Error(), "aborted") {
		log.Fatalln("failed to send message:", err)
	}

	fmt.Println("\n     \033[31m[DISCONNECTED] Lost connection: attempting to reconnect...\033[0m")
	conn.Close()
	conn = getConnection(fmt.Sprintf(":%v", PORT))
	connectionReader.Reset(conn)
	fmt.Printf("     \033[32m[CONNECTED]    Reconnected to the server successfully!\033[0m\n\n")
	return conn, true
}

func main() {
	conn = getConnection(fmt.Sprintf(":%v", PORT))
	defer conn.Close()

	terminalReader := bufio.NewReader(os.Stdin)
	connectionReader := bufio.NewReader(conn)
	for {
		is_reset := false
		fmt.Printf("\033[3;36m>\033[0m \033[1;3;35m")
		message, _ := terminalReader.ReadString('\n')
		fmt.Print("\033[0m")
		_, err := conn.Write([]byte(message))
		conn, is_reset = handleConnectionError(conn, connectionReader, err)
		if is_reset {
			continue
		}

		message, err = connectionReader.ReadString('\n')
		conn, is_reset = handleConnectionError(conn, connectionReader, err)
		if is_reset {
			continue
		}

		fmt.Println("\033[92m\n\tMessage from server: \033[0m", message)
	}
}
