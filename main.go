package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ComradeCaro/SimpleChat/client"
	"github.com/ComradeCaro/SimpleChat/server"
)

func main() {
	fmt.Println("Starting SimpleChat...")

	// Get user input regarding if the user wants to run a server or client
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to run a client or server? C/s: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	input = strings.TrimSpace(input)

	switch strings.ToLower(input) {
	case "c", "client":
		client.Run()
	case "s", "server":
		server.Run()
	default:
		fmt.Println("Defaulting to client.")
		client.Run()
	}
}
