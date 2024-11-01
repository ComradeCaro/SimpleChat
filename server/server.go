package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	clients = make(map[net.Conn]bool) // Map to keep track of connected clients
	mu      sync.Mutex                // Mutex for safe access to clients map
	newPort string
	err     error
)

func Run(port string) {

	if port == "20000" {
		// Asks user what port they want to listen to
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What port do you want to listen on? (Default is 20000): ")
		newPort, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	if newPort != "" {
		port = strings.TrimSpace(newPort)
	} else {
		port = strings.TrimSpace(port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s...\n", port)

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Read the client's name
		nameReader := bufio.NewReader(conn)
		name, err := nameReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading client name:", err)
			conn.Close()
			continue
		}
		name = strings.TrimSpace(name)

		// Add new client to the clients map
		mu.Lock()
		clients[conn] = true
		mu.Unlock()

		// Handle the connection in a new goroutine
		go handleConnection(conn, name)
	}
}

func handleConnection(conn net.Conn, name string) {
	// Send a "joined the chat" message to all clients except the newly connected one
	broadcastMessage(fmt.Sprintf("Client %s has joined the chat!", name), conn)

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()

		// Broadcast that the client has left
		broadcastMessage(fmt.Sprintf("Client %s has left the chat.", name), nil)
		fmt.Printf("Client %v has left the chat.\n", name)
	}()

	fmt.Printf("Client %v connected on: %v\n", name, conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}
		fmt.Printf("%v: %s\n", name, message)

		// Broadcast the message to other clients
		broadcastMessage(fmt.Sprintf("%s: %s", name, message), conn)
	}
}

func broadcastMessage(message string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for clientConn := range clients {
		// Skip the sender client if specified
		if clientConn == sender {
			continue
		}
		if _, err := clientConn.Write([]byte(message + "\n")); err != nil {
			fmt.Println("Error sending broadcast message to client:", err)
		}
	}
}
