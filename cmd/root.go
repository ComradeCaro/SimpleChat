package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ComradeCaro/SimpleChat/client"
	"github.com/ComradeCaro/SimpleChat/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simplechat",
	Short: "SimpleChat is a simple TCP chat application",
	Long:  `SimpleChat allows you to run a server or a client for chatting over TCP.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
