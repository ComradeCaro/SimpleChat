package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Run() {
	// Asks user to enter address to server
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter server address (host:port): ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	address = strings.TrimSpace(address)

	fmt.Print("Please enter a name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	name = strings.TrimSpace(name)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send the name to the server
	_, err = conn.Write([]byte(name + "\n"))
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}

	go readMessages(conn) // Start a goroutine to read messages from the server

	fmt.Printf("Connected to server as %v! Type your messages below:\n", name)
	for {
		// Read input from user
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Exiting client...")
			break
		}

		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Disconnected from server.")
				os.Exit(0)
			} else {
				fmt.Println("Error reading from server:", err)
			}
			return
		}
		fmt.Printf("%s\n> ", strings.TrimSpace(message))
	}
}
