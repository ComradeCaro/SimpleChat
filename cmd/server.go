package cmd

import (
	"github.com/ComradeCaro/SimpleChat/server"
	"github.com/spf13/cobra"
)

var port string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the SimpleChat server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(port, true)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "p", "20000", "Port server will listen on")
}
